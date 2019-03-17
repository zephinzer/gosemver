package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var GitAllTags = []string{"tag", "--list"}
var GitLatestTag = []string{"describe", "--tags", "--abbrev=0"}

func GitLoader() (int, int, int, string, error) {
	verifyGitExists()
	fmt.Println("-----------------------------------------")
	fmt.Println(getAllTags())
	fmt.Println("-----------------------------------------")
	return 0, 0, 0, "?", nil
}

func getAllTags() []string {
	var allTagsOutput bytes.Buffer
	getVersion := exec.Command("git", GitAllTags...)
	getVersion.Stdout = &allTagsOutput
	getVersion.Stderr = &allTagsOutput
	getVersion.Run()

	return strings.Split(
		strings.Replace(
			allTagsOutput.String(),
			"\r",
			"",
			-1,
		),
		"\n",
	)
}

func getLatestVersion() string {
	verifyGitExists()
	var versionOutput bytes.Buffer
	getVersion := exec.Command("git", GitLatestTag...)
	getVersion.Stdout = &versionOutput
	getVersion.Stderr = &versionOutput
	getVersion.Run()

	return versionOutput.String()
}

func verifyGitExists() {
	_, err := exec.LookPath("git")
	if err != nil {
		panic(err)
	}
}
