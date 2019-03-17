package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	semver := &Semver{}
	config := InitialiseConfiguration()
	if config.Use == "git" {
		semver.Load(GitLoader)
	}
	if config.BumpMajor {
		fmt.Println("bump major!")
	} else if config.BumpMinor {
		fmt.Println("bump minor!")
	} else if config.BumpPatch {
		fmt.Println("bump patch!")
	} else if len(config.BumpLabel) > 0 {
		fmt.Println("bump label!")
	} else {
		fmt.Println("?")
		os.Exit(1)
	}
}

// GoSemverConfig is for holding the cli configuration flags
type GoSemverConfig struct {
	BumpMajor bool
	BumpMinor bool
	BumpPatch bool
	BumpLabel string
	Use       string
}

// InitialiseConfiguration initialises the configuration for the cli
func InitialiseConfiguration() *GoSemverConfig {
	config := &GoSemverConfig{}
	flag.BoolVar(&config.BumpMajor, "major", false, "bump the major version")
	flag.BoolVar(&config.BumpMinor, "minor", false, "bump the minor version")
	flag.BoolVar(&config.BumpPatch, "patch", false, "bump the patch version")
	flag.StringVar(&config.BumpLabel, "label", "", "bump the label number")
	flag.StringVar(&config.Use, "use", "git", "indicates the loader to use")
	flag.Parse()
	return config
}
