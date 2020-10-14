package main

import (
	"os"

	"github.com/shuymn/regend/cli"
)

func main() {
	c := cli.NewCLI()
	os.Exit(c.Run(os.Args[1:]))
}
