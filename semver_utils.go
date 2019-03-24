package main

import (
	"sort"
	"strings"
)

// New creates an instance of Semver from scratch
func New(major int, minor int, patch int, label string, prefix ...string) *Semver {
	if len(prefix) > 0 {
		return &Semver{major, minor, patch, label, prefix[0]}
	}
	return &Semver{major, minor, patch, label, ""}
}

// NewFrom creates an instance of Semver given a SemverLoader
func NewFrom(loader SemverLoader) (*Semver, error) {
	semver := &Semver{}
	err := semver.Load(loader)
	return semver, err
}

// Sort is a convenience function for sorting semvers
func Sort(semvers []ISemver) []ISemver {
	sort.Stable(BySemver(semvers))
	return semvers
}

// BySemver is for sorting a slice of Semvers
type BySemver []ISemver

// Len implements the required interface for the `sort` package that returns
// the total length of the slice
func (bySemver BySemver) Len() int {
	return len(bySemver)
}

// Swap implements the required interface for the `sort` package that swaps
// two members of a slice
func (bySemver BySemver) Swap(i, j int) {
	bySemver[i], bySemver[j] = bySemver[j], bySemver[i]
}

// Less implements the required interface for the `sort` package that returns
// true if the element at `i` is less than the element at `j`
func (bySemver BySemver) Less(i, j int) bool {
	return bySemver[i].GetMajorInt() < bySemver[j].GetMajorInt() ||
		(bySemver[i].GetMajorInt() <= bySemver[j].GetMajorInt() &&
			bySemver[i].GetMinorInt() < bySemver[j].GetMinorInt()) ||
		(bySemver[i].GetMajorInt() <= bySemver[j].GetMajorInt() &&
			bySemver[i].GetMinorInt() <= bySemver[j].GetMinorInt() &&
			bySemver[i].GetPatchInt() < bySemver[j].GetPatchInt()) ||
		(bySemver[i].GetMajorInt() <= bySemver[j].GetMajorInt() &&
			bySemver[i].GetMinorInt() <= bySemver[j].GetMinorInt() &&
			bySemver[i].GetPatchInt() <= bySemver[j].GetPatchInt() &&
			bySemver[i].GetLabelPower() > bySemver[j].GetLabelPower()) ||
		(bySemver[i].GetMajorInt() <= bySemver[j].GetMajorInt() &&
			bySemver[i].GetMinorInt() <= bySemver[j].GetMinorInt() &&
			bySemver[i].GetPatchInt() <= bySemver[j].GetPatchInt() &&
			bySemver[i].GetLabelPower() >= bySemver[j].GetLabelPower() &&
			strings.Compare(bySemver[i].GetLabelString(), bySemver[j].GetLabelString()) < 0) ||
		(bySemver[i].GetMajorInt() <= bySemver[j].GetMajorInt() &&
			bySemver[i].GetMinorInt() <= bySemver[j].GetMinorInt() &&
			bySemver[i].GetPatchInt() <= bySemver[j].GetPatchInt() &&
			bySemver[i].GetLabelPower() >= bySemver[j].GetLabelPower() &&
			strings.Compare(bySemver[i].GetLabelString(), bySemver[j].GetLabelString()) <= 0 &&
			bySemver[i].GetLabelInt() < bySemver[j].GetLabelInt())
}
