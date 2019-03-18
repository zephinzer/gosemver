package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func handleBump(c *cli.Context) error {
	switch c.Args().First() {
	case "major":
		fmt.Println("MAJOR BUMP")
		fmt.Println("using: ", c.String("use"))
	case "minor":
		fmt.Println("MINOR BUMP")
		fmt.Println("using: ", c.String("use"))
	case "patch":
		fmt.Println("PATCH BUMP")
		fmt.Println("using: ", c.String("use"))
	case "label":
		fmt.Println("LABEL BUMP - label:" + c.Args().Get(1))
		fmt.Println("using: ", c.String("use"))
	case "":
		fallthrough
	default:
		cli.ShowAppHelp(c)
	}
	return nil
}

func handleDefault(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func handleGet(c *cli.Context) error {
	get := strings.ToLower(c.Args().First())
	using := strings.ToLower(c.String("use"))
	prefix := strings.ToLower(c.String("prefix"))

	var semver ISemver
	var err error
	switch using {
	case "git":
		loader := GitLoader{}
		semver, err = NewFrom(loader.Load("latest", prefix))
	default:
		fmt.Println("invalid engine '" + using + "' specified")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch get {
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

func handleSet(c *cli.Context) error {
	set := strings.ToLower(c.Args().First())
	using := strings.ToLower(c.String("use"))
	prefix := strings.ToLower(c.String("prefix"))

	switch using {
	case "git":
		if isSemverLike(set) {
			semver := toSemver(set, prefix)
			fmt.Println("setting version to:")
			fmt.Printf("  prefix : %s\n", semver.GetPrefix())
			fmt.Printf("  major  : %v\n", semver.GetMajorInt())
			fmt.Printf("  minor  : %v\n", semver.GetMinorInt())
			fmt.Printf("  patch  : %v\n", semver.GetPatchInt())
			fmt.Printf("  label  : %s\n", semver.GetLabel())
			fmt.Printf("  --------\n")
			fmt.Printf("  %s\n", semver)
		} else {
			fmt.Println("invalid semver '" + set + "' specified")
		}
	default:
		fmt.Println("invalid engine '" + using + "' specified")
	}
	return nil
}

func handleVersion(c *cli.Context) error {
	version := strings.ToLower(c.Args().First())

	switch version {
	case "semver":
		fmt.Println(Version)
	case "commit":
		fmt.Println(Commit)
	default:
		cli.ShowVersion(c)
	}
	return nil
}
