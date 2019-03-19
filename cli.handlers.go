package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func handleBump(c *cli.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	section := c.Args().First()
	using := c.String("use")
	label := c.Args().Get(1)
	if err := cliBump(section, using, label); err != nil {
		cli.ShowAppHelp(c)
		return err
	}
	return nil
}

func handleDefault(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func handleGet(c *cli.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	section := strings.ToLower(c.Args().First())
	using := strings.ToLower(c.String("use"))
	prefix := strings.ToLower(c.String("prefix"))

	return cliGet(section, using, prefix)
}

func handleSet(c *cli.Context) error {
	version := strings.ToLower(c.Args().First())
	using := strings.ToLower(c.String("use"))
	prefix := strings.ToLower(c.String("prefix"))

	return cliSet(version, using, prefix)
}

func handleVersion(c *cli.Context) error {
	version := strings.ToLower(c.Args().First())
	if err := cliVersion(version); err != nil {
		cli.ShowVersion(c)
		return err
	}
	return nil
}
