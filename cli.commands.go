package main

import "github.com/urfave/cli"

type commandProvider func() cli.Command

func commands(commands ...commandProvider) []cli.Command {
	var commandChain []cli.Command
	for _, command := range commands {
		commandChain = append(commandChain, command())
	}
	return commandChain
}

func commandBump() cli.Command {
	return cli.Command{
		Action:      handleBump,
		Aliases:     []string{"b"},
		ArgsUsage:   "<major, minor, patch, label>",
		Description: "if no arguments are specified, defaults to bumping the patch version.",
		Flags:       flags(flagUse),
		Name:        "bump",
		Usage:       "bump a version",
	}
}

func commandGet() cli.Command {
	return cli.Command{
		Action:      handleGet,
		Aliases:     []string{"g"},
		ArgsUsage:   "<major, minor, patch, label>",
		Description: "gets the version of the application under development - to retrieve a specific section, use one of 'major', 'minor', 'patch', 'label', otherwise the entire version will be returned if not arguments are specified.",
		Flags:       flags(flagUse, flagPrefix, flagMode),
		Name:        "get",
	}
}

func commandSet() cli.Command {
	return cli.Command{
		Action:      handleSet,
		Aliases:     []string{"s"},
		ArgsUsage:   "<< version to set >>",
		Description: "sets the version of the application under development to a specific version of your choice.",
		Flags:       flags(flagUse),
		Name:        "set",
	}
}

func commandVersion() cli.Command {
	return cli.Command{
		Action:      handleVersion,
		Aliases:     []string{"v"},
		ArgsUsage:   "<semver, commit>",
		Description: "use 'semver' to retrieve only the X.Y.Z version, or 'commit' to retrieve the commit hash. defaults to retrieving the full version of <VERSION>-<COMMIT>",
		Name:        "version",
		Usage:       "retrieves gosemver's version",
	}
}
