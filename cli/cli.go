package cli

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/shuymn/regen/config"
	"github.com/shuymn/regen/static"
)

const name = "regen"

const (
	exitCodeOK = iota
	exitCodeErr
)

type CLI struct {
	conf *config.RegenConfig
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) Run(args []string) int {
	conf := config.NewConfig()

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

	tomlFile := config.LoadTOMLFilename()
	if tomlFile != "" {
		if err := conf.LoadTOML(tomlFile); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
			return exitCodeErr
		}
	}
	c.conf = &conf.Regen

	// environmennnt variables
	if regenHost := os.Getenv("REGEN_HOST"); regenHost != "" {
		c.conf.Host = regenHost
	}
	regenPort, err := strconv.Atoi(os.Getenv("REGEN_PORT"))
	if err == nil && regenPort != 0 {
		c.conf.Port = regenPort
	}
	if regenUser := os.Getenv("REGEN_USER"); regenUser != "" {
		c.conf.User = regenUser
	}
	if regenPassword := os.Getenv("REGEN_PASSWORD"); regenPassword != "" {
		c.conf.Password = regenPassword
	}
	if regenDatabase := os.Getenv("REGEN_DATABASE"); regenDatabase != "" {
		c.conf.Database = regenDatabase
	}

	if c.conf.Host == "" {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, "must specify host")
		return exitCodeErr
	}
	if c.conf.Port == 0 {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, "must specify port")
		return exitCodeErr
	}
	if c.conf.User == "" {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, "must specify user")
		return exitCodeErr
	}
	if c.conf.Password == "" {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, "must specify password")
		return exitCodeErr
	}
	if c.conf.Database == "" {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, "must specify database")
		return exitCodeErr
	}

	if err := c.generate(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		return exitCodeErr
	}
	return exitCodeOK
}

func (c *CLI) generate(table string) error {
	driverName := "postgres"
	connStr := fmt.Sprintf("%s://%s:%s@%s:%d/%s", driverName, c.conf.User, c.conf.Password, c.conf.Host, c.conf.Port, c.conf.Database)
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
