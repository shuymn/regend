package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	"github.com/shuymn/regen/static"
)

const name = "regen"

const (
	exitCodeOK = iota
	exitCodeErr
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return exitCodeOK
		}
		return exitCodeErr
	}
	args = fs.Args()
	if len(args) == 0 || len(args) > 1 {
		fmt.Fprintf(os.Stderr, "usage: %s table\n", name)
		return exitCodeErr
	}
	if err := generate(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		return exitCodeErr
	}
	return exitCodeOK
}

func generate(table string) error {
	driverName := "postgres"
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s", driverName, os.Getenv("REDSHIFT_USER"), os.Getenv("REDSHIFT_PASSWORD"), os.Getenv("REDSHIFT_HOST"), os.Getenv("REDSHIFT_PORT"), os.Getenv("REDSHIFT_DATABASE"))
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	file, err := static.Root.Open("/generate_tbl_ddl.sql")
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(content)

	rows, err := db.Query(b.String(), table)
	if err != nil {
		return err
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		for i, col := range values {
			if columns[i] == "ddl" {
				fmt.Println(string(col))
			}
		}
	}

	return nil
}
