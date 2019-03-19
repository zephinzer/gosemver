package main

import (
	"errors"
	"fmt"
	"os"
)

func cliBump(section string, using string, anyLabels ...string) error {
	var label string
	if len(anyLabels) > 0 {
		label = anyLabels[0]
	}
	switch section {
	case "major":
		fmt.Println("MAJOR BUMP")
		fmt.Println("using: ", using)
	case "minor":
		fmt.Println("MINOR BUMP")
		fmt.Println("using: ", using)
	case "label":
		fmt.Println("LABEL BUMP - label:" + label)
		fmt.Println("using: ", using)
	case "":
		fallthrough
	case "patch":
		fallthrough
	default:
		fmt.Println("PATCH BUMP")
		fmt.Println("using: ", using)
	}
	return nil
}

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

func cliSet(version string, using string, prefix string) error {
	switch using {
	case "git":
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
	default:
		fmt.Println("invalid engine '" + using + "' specified")
	}
	return nil
}

func cliVersion(version string) error {
	switch version {
	case "semver":
		fmt.Println(Version)
	case "commit":
		fmt.Println(Commit)
	default:
		return errors.New("invalid version type '" + version + "' specified")
	}
	return nil
}
