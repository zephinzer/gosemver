package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

var GitAllTags = []string{"tag", "--list"}
var GitLatestTag = []string{"describe", "--tags", "--abbrev=0"}

func NewGitLoader(prefix ...string) func() (int, int, int, string, error) {
	return func() (int, int, int, string, error) {
		verifyGitExists()
		allTags := getAllTags()
		allSemverTags := filterSemverLike(allTags, prefix...)
		var semvers []ISemver
		var major int
		var minor int
		var patch int
		var label string
		var err error
		for _, semverTag := range allSemverTags {
			semverLabels := strings.Split(semverTag, "-")
			semverSections := strings.Split(semverLabels[0], ".")
			if major, err = strconv.Atoi(semverSections[0]); err != nil {
				panic(err)
			}
			if minor, err = strconv.Atoi(semverSections[1]); err != nil {
				panic(err)
			}
			if patch, err = strconv.Atoi(semverSections[2]); err != nil {
				panic(err)
			}
			if len(semverLabels) > 1 {
				label = strings.Join(semverLabels[1:], "-")
			} else {
				label = ""
			}
			semvers = append(semvers, &Semver{major, minor, patch, label})
		}
		semvers = Sort(semvers)
		latest := semvers[len(semvers)-1]
		return latest.GetMajorInt(), latest.GetMinorInt(), latest.GetPatchInt(), latest.GetLabel(), nil
	}
}

func getAllTags() []string {
	var allTagsRawOutput bytes.Buffer
	getVersion := exec.Command("git", GitAllTags...)
	getVersion.Stdout = &allTagsRawOutput
	getVersion.Stderr = &allTagsRawOutput
	getVersion.Run()

	tagsList := strings.Split(
		strings.Replace(
			allTagsRawOutput.String(),
			"\r",
			"",
			-1,
		),
		"\n",
	)

	var allTags []string
	for _, tag := range tagsList {
		if len(tag) > 0 {
			allTags = append(allTags, tag)
		}
	}

	return allTags
}

func getLatestVersion() string {
	verifyGitExists()
	var versionOutput bytes.Buffer
	getVersion := exec.Command("git", GitLatestTag...)
	getVersion.Stdout = &versionOutput
	getVersion.Stderr = &versionOutput
	getVersion.Run()

	return strings.Trim(strings.Replace(versionOutput.String(), "\r", "", -1), "\r\n ")
}

// verifyGitExists panics if Git is not found
func verifyGitExists() {
	_, err := exec.LookPath("git")
	if err != nil {
		panic(err)
	}
}
