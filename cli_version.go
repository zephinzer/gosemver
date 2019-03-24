package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

type CLIVersion func(string) error

func cliVersion(version string) error {
	switch version {
	case "help":
		return fmt.Errorf("help requested")
	case "semver":
		fmt.Println(Version)
	case "commit":
		fmt.Println(Commit)
	default:
		fmt.Printf("godev %s-%s\n", Version, Commit)
	}
	return nil
}

func handleVersion(c *cli.Context, version CLIVersion) error {
	versionType := strings.ToLower(c.Args().First())
	if err := version(versionType); err != nil {
		cli.ShowSubcommandHelp(c)
		return err
	}
	return nil
}

func getVersionCommand() cli.Command {
	return cli.Command{
		Action: func(c *cli.Context) {
			handleVersion(c, cliVersion)
		},
		Aliases:     []string{"v"},
		ArgsUsage:   "<< semver | commit >>",
		Description: "use 'semver' to retrieve only the X.Y.Z version, or 'commit' to retrieve the commit hash. defaults to retrieving the full version of 'godev <<VERSION>>-<<COMMIT>>'",
		Name:        "version",
		Usage:       "retrieve gosemver's version",
	}
}
