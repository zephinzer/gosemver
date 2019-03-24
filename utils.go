package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

// commandProvider defines a function that returns a command handler
type commandProvider func() cli.Command

// flagProvider defines a function that returns a flag handler
type flagProvider func() cli.Flag

// confirmationFalseCanonical defines the canonical rejection character
const confirmationFalseCanonical = "n"

// confirmationFalseCanonical defines the canonical acceptance character
const confirmationTrueCanonical = "y"

// semverIntSection for use in regexp building
const semverIntSection = `(0|[1-9]{1}[\d]*){1}`

// semverIntSeperator for use in regexp building
const semverIntSeperator = `\.`

// semverLabelSection for use in regexp building
const semverLabelSection = `(\-[a-zA-Z0-9\.\_]*)*`

// confirmationFalse defines the alises of a rejection
var confirmationFalse = []string{confirmationFalseCanonical, "no", "nope", "nah", "neh", "stop", "dont"}

// confirmationTrue defines the alises of an acceptance
var confirmationTrue = []string{confirmationTrueCanonical, "yes", "yupp", "yeah", "yea", "ok", "okay"}

// commands returns a slice of commands that can be used by a
// command orchestrator
func commands(commands ...commandProvider) []cli.Command {
	var commandChain []cli.Command
	for _, command := range commands {
		commandChain = append(commandChain, command())
	}
	return commandChain
}

func confirm(reader *bufio.Reader, question string, byDefault bool, retryText ...string) bool {
	var options string
	if byDefault {
		options = fmt.Sprintf("%s/%s", strings.ToUpper(confirmationTrueCanonical), confirmationFalseCanonical)
	} else {
		options = fmt.Sprintf("%s/%s", confirmationTrueCanonical, strings.ToUpper(confirmationFalseCanonical))
	}
	fmt.Printf("%s [%s]: ", question, options)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if len(userInput) < 2 {
		return byDefault
	}
	content := strings.Trim(
		strings.ToLower(userInput),
		" \r\n.,;",
	)
	confirmation := false
	if sliceContainsString(confirmationTrue, content) {
		confirmation = true
	} else if sliceContainsString(confirmationFalse, content) {
		confirmation = false
	} else if len(retryText) > 0 {
		fmt.Println(retryText[0])
		confirmation = confirm(reader, question, byDefault, retryText...)
	}
	return confirmation
}

// flags returns a slice of flags that can be used in a command
func flags(flags ...flagProvider) []cli.Flag {
	var flagChain []cli.Flag
	for _, flag := range flags {
		flagChain = append(flagChain, flag())
	}
	return flagChain
}

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

// sliceContainsString returns true if the :slice contains the :search
// string
func sliceContainsString(slice []string, search string) bool {
	for _, sliceItem := range slice {
		if search == sliceItem {
			return true
		}
	}
	return false
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
