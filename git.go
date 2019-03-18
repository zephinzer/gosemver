package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type GitLoader struct{}

func (gitLoader *GitLoader) Load(mode string, prefix ...string) SemverLoader {
	verifyGitExists()
	return func() (int, int, int, string, string, error) {
		var latest ISemver
		if mode == "latest" {
			latest = gitLoader.getLatest(prefix...)
		} else if mode == "current" {
			latest = gitLoader.getCurrent(prefix...)
		}
		if latest == nil {
			return 0, 0, 0, "", "", errors.New("no tags found")
		}
		return latest.GetMajorInt(),
			latest.GetMinorInt(),
			latest.GetPatchInt(),
			latest.GetLabel(),
			latest.GetPrefix(),
			nil
	}
}

func (gitLoader *GitLoader) getLatest(prefix ...string) ISemver {
	semvers := make([]ISemver, 0)
	semverStrings := filterSemverLike(
		gitLoader.getAllTags(prefix...),
		prefix...,
	)
	for _, semverTag := range semverStrings {
		semvers = append(semvers, toSemver(semverTag, prefix...))
	}
	semvers = Sort(semvers)
	return semvers[len(semvers)-1]
}

func (gitLoader *GitLoader) getCurrent(prefix ...string) ISemver {
	return toSemver(gitDescribeTag(), prefix...)
}

func (gitLoader *GitLoader) getAllTags(prefix ...string) []string {
	return removeEmptyStringsFromStringSlice(strings.Split(gitTagList(), "\n"))
}

// gitTagList retrieves all git tags
func gitTagList() string {
	var allTagsRawOutput bytes.Buffer
	getVersion := exec.Command("git", "tag", "--list")
	getVersion.Stdout = &allTagsRawOutput
	getVersion.Stderr = &allTagsRawOutput
	getVersion.Run()
	return trimAndNormalise(allTagsRawOutput.String())
}

// gitDescribeTag retrieves the most recent tag
func gitDescribeTag() string {
	var versionOutput bytes.Buffer
	getVersion := exec.Command("git", "describe", "--tags", "--abbrev=0")
	getVersion.Stdout = &versionOutput
	getVersion.Stderr = &versionOutput
	getVersion.Run()
	return trimAndNormalise(versionOutput.String())
}

// gitTag tags a commit with the given tag
func gitTag(tag string) string {
	var anyOutput bytes.Buffer
	getVersion := exec.Command("git", "tag", tag)
	getVersion.Stdout = &anyOutput
	getVersion.Stderr = &anyOutput
	getVersion.Run()
	return trimAndNormalise(anyOutput.String())
}

// verifyGitExists panics if Git is not found
func verifyGitExists() {
	if _, err := exec.LookPath("git"); err != nil {
		panic("the git executable could not be found/is not in your path")
	}
}
