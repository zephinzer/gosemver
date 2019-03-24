package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

type CLISet func(string, string) error

func cliSet(version string, prefix string) error {
	if version == "help" {
		return fmt.Errorf("help requested")
	}
	if isSemverLike(version) {
		semver := toSemver(version, prefix)
		fmt.Println("setting version to:")
		fmt.Printf("  prefix : %s\n", semver.GetPrefix())
		fmt.Printf("  major  : %v\n", semver.GetMajorInt())
		fmt.Printf("  minor  : %v\n", semver.GetMinorInt())
		fmt.Printf("  patch  : %v\n", semver.GetPatchInt())
		fmt.Printf("  label  : %s\n", semver.GetLabel())
		fmt.Printf("  --------\n")
		fmt.Printf("  %s\n", semver)
	} else {
		fmt.Println("invalid semver '" + version + "' specified")
	}
	return nil
}

func getSetCommand() cli.Command {
	return cli.Command{
		Action: func(c *cli.Context) {
			handleSet(c, cliSet)
		},
		Aliases:     []string{"s"},
		ArgsUsage:   "<< version to set >>",
		Description: "sets the version of the application under development to a specific version of your choice",
		Flags:       flags(flagPrefix),
		Name:        "set",
		Usage:       "explicitly sets the version",
	}
}

func handleSet(c *cli.Context, set CLISet) error {
	version := strings.ToLower(c.Args().First())
	prefix := strings.ToLower(c.String("prefix"))

	if err := set(version, prefix); err != nil {
		cli.ShowSubcommandHelp(c)
		return err
	}
	return nil
}
