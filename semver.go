package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ISemver defines an interface for using Semver
type ISemver interface {
	BumpMajor()
	BumpMinor()
	BumpPatch()
	BumpLabel(string)
	GetMajorInt() int
	GetMajorString() string
	GetMinorInt() int
	GetMinorString() string
	GetPatchInt() int
	GetPatchString() string
	GetLabel() string
	Load(SemverLoader)
	String() string
}

// SemverLoader should load a Semver value from an arbitrary source
type SemverLoader func() (int, int, int, string, error)

// New creates an instance of Semver from scratch
func New(major int, minor int, patch int, label string) *Semver {
	return &Semver{major, minor, patch, label}
}

// NewFrom creates an instance of Semver given a SemverLoader
func NewFrom(loader SemverLoader) *Semver {
	semver := &Semver{}
	semver.Load(loader)
	return semver
}

// Semver holds the data structure for a semantic versioning model
type Semver struct {
	major int
	minor int
	patch int
	label string
}

// BumpMajor adds 1 to the major version and resets the minor and patch version to 0
func (semver *Semver) BumpMajor() {
	semver.major++
	semver.minor = 0
	semver.patch = 0
	semver.label = ""
}

// BumpMinor adds 1 to the minor version and resets the patch version to 0
func (semver *Semver) BumpMinor() {
	semver.minor++
	semver.patch = 0
	semver.label = ""
}

// BumpPatch adds 1 to the patch version
func (semver *Semver) BumpPatch() {
	semver.patch++
	semver.label = ""
}

// BumpLabel checks if the current label is present, if it is, it bumps the
// last number set by one, otherwise, it sets the label and appends a `.0` to
// the label
func (semver *Semver) BumpLabel(label string) {
	if strings.Index(semver.label, label) == 0 {
		existingLabel := strings.Split(semver.label, ".")
		if len(existingLabel) == 1 { // label
			semver.label = semver.label + ".0"
		} else if len(existingLabel) > 2 { // label.rc.X
			stringSection := strings.Join(existingLabel[:len(existingLabel)-1], ".")
			labelVersion, err := strconv.Atoi(existingLabel[len(existingLabel)-1])
			if err != nil {
				panic(err)
			}
			semver.label = fmt.Sprintf("%s.%v", stringSection, labelVersion+1)
		} else { // label.Y
			stringSection := existingLabel[0]
			labelVersion, err := strconv.Atoi(existingLabel[1])
			if err != nil {
				panic(err)
			}
			semver.label = fmt.Sprintf("%s.%v", stringSection, labelVersion+1)
		}
	} else {
		semver.label = label + ".0"
	}
}

// GetMajorInt retrieves the major version number
func (semver *Semver) GetMajorInt() int {
	return semver.major
}

// GetMajorString retrieves the major version number as a string
func (semver *Semver) GetMajorString() string {
	return strconv.Itoa(semver.major)
}

// GetMinorInt retrieves the minor version number
func (semver *Semver) GetMinorInt() int {
	return semver.minor
}

// GetMinorString retrieves the minor version number as a string
func (semver *Semver) GetMinorString() string {
	return strconv.Itoa(semver.minor)
}

// GetPatchInt retrieves the patch version number
func (semver *Semver) GetPatchInt() int {
	return semver.patch
}

// GetPatchString retrieves the patch version number as a string
func (semver *Semver) GetPatchString() string {
	return strconv.Itoa(semver.patch)
}

// GetLabel retrieves the label section
func (semver *Semver) GetLabel() string {
	return semver.label
}

func (semver *Semver) Load(from SemverLoader) {
	major, minor, patch, label, err := from()
	if err != nil {
		panic(err)
	}
	semver.major = major
	semver.minor = minor
	semver.patch = patch
	semver.label = label
}

func (semver *Semver) String() string {
	version := fmt.Sprintf("%v.%v.%v", semver.major, semver.minor, semver.patch)
	if len(semver.label) > 0 {
		version = fmt.Sprintf("%s-%s", version, semver.label)
	}
	return version
}
