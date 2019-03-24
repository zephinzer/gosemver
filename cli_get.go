package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type CLIGet func(string, string, string) error

func cliGet(section string, using string, prefix string) error {
	var semver ISemver
	var err error
	switch using {
	case "git":
		loader := GitLoader{}
		semver, err = NewFrom(loader.Load("latest", prefix))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("invalid engine '" + using + "' specified")
		os.Exit(1)
	}

	switch section {
	case "help":
		return fmt.Errorf("help requested")
	case "major":
		fmt.Println(semver.GetMajorInt())
	case "minor":
		fmt.Println(semver.GetMinorInt())
	case "patch":
		fmt.Println(semver.GetPatchInt())
	case "label":
		fmt.Println(semver.GetLabel())
	default:
		fmt.Println(semver)
	}

	return nil
}

func getGetCommand() cli.Command {
	return cli.Command{
		Action: func(c *cli.Context) {
			handleGet(c, cliGet)
		},
		Aliases:     []string{"g"},
		ArgsUsage:   "<< major | minor | patch | label >>",
		Description: "gets the version of the application under development - to retrieve a specific section, use one of 'major', 'minor', 'patch', 'label', otherwise the entire version will be returned if no arguments are specified.",
		Flags:       flags(flagUse, flagPrefix, flagMode),
		Name:        "get",
		Usage:       "gets the repository's latest/highest tag",
	}
}

func handleGet(c *cli.Context, get CLIGet) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	section := strings.ToLower(c.Args().First())
	using := strings.ToLower(c.String("use"))
	prefix := strings.ToLower(c.String("prefix"))

	if err := get(section, using, prefix); err != nil {
		cli.ShowSubcommandHelp(c)
		return err
	}
	return nil
}
