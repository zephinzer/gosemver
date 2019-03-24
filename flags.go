package main

import "github.com/urfave/cli"

func flagUse() cli.Flag {
	return cli.StringFlag{
		Usage:  "currently it's only 'git'",
		Name:   "use, u",
		Value:  "git",
		EnvVar: "USE",
	}
}

func flagMode() cli.Flag {
	return cli.StringFlag{
		Usage:  "one of 'latest' or 'current': 'latest' gets the latest version, 'current' gets the most recently tagged semver version",
		Name:   "mode, m",
		Value:  "latest",
		EnvVar: "MODE",
	}
}

func flagPrefix() cli.Flag {
	return cli.StringFlag{
		Usage:  "set this to 'v' if your versions are prefixed with a 'v' (eg. v1.0.0)",
		Name:   "prefix, p",
		Value:  "",
		EnvVar: "PREFIX",
	}
}

func flagYes() cli.Flag {
	return cli.BoolFlag{
		Usage:  "specify this to say yes to any questions asked in interactive mode",
		Name:   "yes, y",
		EnvVar: "YES",
	}
}
