package main

import (
	"regexp"
	"strconv"
	"strings"
)

// semverIntSection for use in regexp building
const semverIntSection = `(0|[1-9]{1}[\d]*){1}`

// semverIntSeperator for use in regexp building
const semverIntSeperator = `\.`

// semverLabelSection for use in regexp building
const semverLabelSection = `(\-[a-zA-Z0-9\.\_]*)*`

// filterSemverLike retrieves all semver variants from the :versionList
// parameter
func filterSemverLike(versionList []string, prefix ...string) []string {
	var semverVersionList []string
	for _, version := range versionList {
		if isSemverLike(version, prefix...) {
			semverVersionList = append(semverVersionList, version)
		}
	}
	return semverVersionList
}

// isSemverLike returns true if the provided :version follows a semver
// pattern, and false otherwise
func isSemverLike(version string, prefix ...string) bool {
	prefixString := ""
	if len(prefix) > 0 {
		prefixString = prefixString + prefix[0]
	}
	regexpString :=
		"^" + // start
			"(" + prefixString + `)` + // for prefix strings
			semverIntSection + // for major version
			semverIntSeperator + // .
			semverIntSection + // for minor version
			semverIntSeperator + // .
			semverIntSection + // for patch versoin
			semverLabelSection + // for label string (-<LABEL>)
			"$" // end
	if matched, err := regexp.Match(regexpString, []byte(version)); err != nil {
		panic(err)
	} else if matched {
		return true
	}
	return false
}

func removeEmptyStringsFromStringSlice(slice []string) []string {
	var finalSlice []string
	for _, sliceItem := range slice {
		if len(sliceItem) > 0 {
			finalSlice = append(finalSlice, sliceItem)
		}
	}
	return finalSlice
}

// toSemver converts the string :from into the Semver struct
func toSemver(from string, prefix ...string) ISemver {
	major := 0
	minor := 0
	patch := 0
	label := ""
	semverSections := strings.Split(from, "-")
	versionPrefix := ""
	if len(prefix) > 0 {
		versionPrefix = prefix[0]
	}
	semver := strings.Trim(semverSections[0], versionPrefix)
	semverNumbers := strings.Split(semver, ".")
	major, _ = strconv.Atoi(semverNumbers[0])
	minor, _ = strconv.Atoi(semverNumbers[1])
	patch, _ = strconv.Atoi(semverNumbers[2])
	if len(semverSections) > 1 {
		label = strings.Join(semverSections[1:], "-")
	}
	return &Semver{major, minor, patch, label, versionPrefix}
}

// trimAndNormalise is for making sure we're Windows compatible
func trimAndNormalise(value string) string {
	return strings.Trim(
		strings.Replace(
			value,
			"\r",
			"",
			-1,
		),
		"\n ",
	)
}
