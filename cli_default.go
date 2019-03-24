package main

import (
	"github.com/urfave/cli"
)

func actionDefault(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}
