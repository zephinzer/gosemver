package main

import (
	"fmt"
	"regexp"
	"sort"
)

// filterSemverLike retrieves all semver variants from the :versionList
// parameter
func filterSemverLike(versionList []string, prefix ...string) []string {
	var semverVersionList []string
	prefixString := ""
	if len(prefix) > 0 {
		prefixString = prefixString + prefix[0]
	}
	regexpString := "^(" + prefixString + `)[\d]+\.[\d]+\.[\d]+`
	for _, version := range versionList {
		if matched, err := regexp.Match(regexpString, []byte(version)); err != nil {
			panic(err)
		} else if matched {
			semverVersionList = append(semverVersionList, version)
		}
	}
	return semverVersionList
}

// sortSemver sorts a given list of semver strings optionally with a prefix
// - define the :prefix paramter to indicate that the semver versions are prefixed
// with a sequence of characters (eg. v1.0.3)
func sortSemver(versionList []ISemver) []ISemver {
	fmt.Println(versionList)
	sort.SliceStable(versionList, getSemverSort(versionList))
	fmt.Println(versionList)
	return versionList
}

func getSemverSort(versionList []ISemver) func(i, j int) bool {
	return func(i, j int) bool {
		// if versionList[i].GetMajorInt() > versionList[j].GetMajorInt() {
		// 	return true
		// } else if versionList[i].GetMinorInt() > versionList[j].GetMinorInt() {
		// 	return true
		// } else if versionList[i].GetPatchInt() > versionList[j].GetPatchInt() {
		// 	return true
		// }
		return false
	}
}
