package main

import (
	"os"

	"github.com/shuymn/regen/cli"
)

func main() {
	c := cli.NewCLI()
	os.Exit(c.Run(os.Args[1:]))
}
