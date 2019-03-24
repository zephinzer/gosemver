package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli"
)

type CLIBump func(string, bool, string, ...string) error

func cliBump(bumpType string, ciMode bool, prefix string, anyLabels ...string) error {
	var label string
	if len(anyLabels) > 0 {
		label = anyLabels[0]
	}
	loader := GitLoader{}
	semver, err := NewFrom(loader.Load("latest", prefix))
	if err != nil {
		panic(err)
	}
	currentSemver := semver.String()
	confirmed := ciMode
	switch bumpType {
	case "help":
		return fmt.Errorf("help requested")
	case "major":
		semver.BumpMajor()
	case "minor":
		semver.BumpMinor()
	case "label":
		semver.BumpLabel(label)
	case "":
		fallthrough
	case "patch":
		fallthrough
	default:
		semver.BumpPatch()
		bumpType = "patch"
	}
	nextSemver := semver.String()
	if !confirmed {
		confirmed = bumpConfirm(os.Stdin, bumpType, currentSemver, nextSemver)
	}
	if confirmed {
		gitTag(nextSemver)
	}
	return nil
}

func getBumpCommand() cli.Command {
	return cli.Command{
		Action: func(c *cli.Context) {
			handleBump(c, cliBump)
		},
		Aliases:     []string{"b"},
		ArgsUsage:   "<< major | minor | patch | label >>",
		Description: "bumps the repositories version. if no arguments are specified, defaults to bumping the patch version",
		Flags:       flags(flagPrefix, flagYes),
		Name:        "bump",
		Usage:       "bumps the repository's version",
	}
}

func handleBump(c *cli.Context, bump CLIBump) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	section := c.Args().First()
	prefix := c.String("prefix")
	label := c.Args().Get(1)
	yes := c.Bool("yes")
	if err := bump(section, yes, prefix, label); err != nil {
		cli.ShowSubcommandHelp(c)
		return err
	}
	return nil
}

func bumpConfirm(via io.Reader, bumpType string, preBump string, postBump string) bool {
	return confirm(
		bufio.NewReader(via),
		fmt.Sprintf(
			"bump the %s version (%s -> %s)? ",
			bumpType,
			preBump,
			postBump,
		),
		false,
	)
}
