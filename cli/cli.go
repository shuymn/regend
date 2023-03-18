package cli

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/shuymn/regend/config"
	"github.com/shuymn/regend/static"
)

const name = "regend"

const (
	exitCodeOK = iota
	exitCodeErr
)

type CLI struct {
	conf *config.RedshiftConfig
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
	c.conf = &conf.Redshift

	// environmennnt variables
	if redshiftHost := os.Getenv("REGEND_REDSHIFT_HOST"); redshiftHost != "" {
		c.conf.Host = redshiftHost
	}
	redshiftPort, err := strconv.Atoi(os.Getenv("REGEND_REDSHIFT_PORT"))
	if err == nil && redshiftPort != 0 {
		c.conf.Port = redshiftPort
	}
	if redshiftUser := os.Getenv("REGEND_REDSHIFT_USER"); redshiftUser != "" {
		c.conf.User = redshiftUser
	}
	if redshiftPassword := os.Getenv("REGEND_REDSHIFT_PASSWORD"); redshiftPassword != "" {
		c.conf.Password = redshiftPassword
	}
	if redshiftDatabase := os.Getenv("REGEND_REDSHIFT_DATABASE"); redshiftDatabase != "" {
		c.conf.Database = redshiftDatabase
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

	content, err := io.ReadAll(file)
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
