package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ISemver defines an interface for using Semver
type ISemver interface {
	BumpLabel(string)
	BumpMajor()
	BumpMinor()
	BumpPatch()
	GetLabel() string
	GetLabelInt() int
	GetLabelString() string
	GetLabelPower() int
	GetMajorInt() int
	GetMinorInt() int
	GetPatchInt() int
	GetPrefix() string
	Load(SemverLoader) error
	String() string
}

// SemverLoader should load a Semver value from an arbitrary source
type SemverLoader func() (int, int, int, string, string, error)

// Semver holds the data structure for a semantic versioning model
type Semver struct {
	major  int
	minor  int
	patch  int
	label  string
	prefix string
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

// GetMinorInt retrieves the minor version number
func (semver *Semver) GetMinorInt() int {
	return semver.minor
}

// GetPatchInt retrieves the patch version number
func (semver *Semver) GetPatchInt() int {
	return semver.patch
}

// GetLabel retrieves the entire label as-is
func (semver *Semver) GetLabel() string {
	return semver.label
}

// GetLabelInt retrieves the label version number. If there is no version number,
// it is taken that the label is the final version of the label series and is
// assigned a maximum integer value for sorting purposes
func (semver *Semver) GetLabelInt() int {
	semverSections := strings.Split(semver.label, ".")
	if len(semverSections) < 2 {
		return math.MaxInt32
	} else {
		semverSectionsLastElement := semverSections[len(semverSections)-1]
		if labelVersion, err := strconv.Atoi(semverSectionsLastElement); err != nil {
			return math.MaxInt32
		} else {
			return labelVersion
		}
	}
}

// GetLabelString retrieves the string section of the label value
func (semver *Semver) GetLabelString() string {
	semverSections := strings.Split(semver.label, ".")
	if len(semverSections) == 0 {
		return ""
	} else if len(semverSections) == 1 {
		return semverSections[0]
	} else {
		semverSectionsLastElement := semverSections[len(semverSections)-1]
		if _, err := strconv.Atoi(semverSectionsLastElement); err != nil {
			return semver.label
		} else {
			return strings.Join(semverSections[:len(semverSections)-1], ".")
		}
	}
}

// GetLabelPower retrieves the power of the label - a lower power indicates
// higher order
func (semver *Semver) GetLabelPower() int {
	labelString := semver.GetLabelString()
	labelSections := strings.Split(labelString, ".")
	labelPower := len(labelSections)
	if labelPower == 1 && len(labelString) == 0 {
		labelPower = 0
	}
	return labelPower
}

// GetPrefix retrieves the prefix of the semver version
func (semver *Semver) GetPrefix() string {
	return semver.prefix
}

// Load loads the Semver struct with values from a loader function implementing
// the SemverLoader type
func (semver *Semver) Load(from SemverLoader) error {
	major, minor, patch, label, prefix, err := from()
	if err != nil {
		return err
	}
	semver.major = major
	semver.minor = minor
	semver.patch = patch
	semver.label = label
	semver.prefix = prefix
	return nil
}

// String converts the Semver struct into its string representation
func (semver *Semver) String() string {
	version := fmt.Sprintf("%s%v.%v.%v", semver.prefix, semver.major, semver.minor, semver.patch)
	if len(semver.label) > 0 {
		version = fmt.Sprintf("%s-%s", version, semver.label)
	}
	return version
}
