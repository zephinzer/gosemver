package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
	testTagList []string
}

func TestUtils(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

func (s *UtilsTestSuite) SetupTest() {
	s.testTagList = []string{
		"0.0.0", // base case
		"v0.0.0",
		"ver-0.0.0",
		"0.0.1",
		"v0.0.1",
		"ver-0.0.1",
		"0.0.9",
		"v0.0.9",
		"ver-0.0.9",
		"non-semver", // dumbest fail case
		"0.1.0",
		"v0.1.0",
		"ver-0.1.0",
		"0.9.0",
		"v0.9.0",
		"ver-0.9.0",
		"a.b.c", // alphabetical fail case
		"1.0.0",
		"v1.0.0",
		"ver-1.0.0",
		"9.0.0",
		"v9.0.0",
		"ver-9.0.0",
		"1.0",     // incomplete fail case #1
		"v1.0",    // incomplete fail case #1
		"ver-1.0", // incomplete fail case #1
		"1",       // incomplete fail case #2
		"v1",      // incomplete fail case #2
		"ver-1",   // incomplete fail case #2
	}
}

func (s *UtilsTestSuite) Test_filterSemverLikeWithoutFilter() {
	assert.Equal(
		s.T(),
		[]string{
			"0.0.0",
			"0.0.1",
			"0.0.9",
			"0.1.0",
			"0.9.0",
			"1.0.0",
			"9.0.0",
		},
		filterSemverLike(s.testTagList),
	)
}

func (s *UtilsTestSuite) Test_filterSemverLikeWithFilter() {
	assert.Equal(
		s.T(),
		[]string{
			"v0.0.0",
			"v0.0.1",
			"v0.0.9",
			"v0.1.0",
			"v0.9.0",
			"v1.0.0",
			"v9.0.0",
		},
		filterSemverLike(s.testTagList, "v"),
	)
}

func (s *UtilsTestSuite) Test_filterSemverLikeWithMutliCharFilter() {
	assert.Equal(
		s.T(),
		[]string{
			"ver-0.0.0",
			"ver-0.0.1",
			"ver-0.0.9",
			"ver-0.1.0",
			"ver-0.9.0",
			"ver-1.0.0",
			"ver-9.0.0",
		},
		filterSemverLike(s.testTagList, "ver-"),
	)
}

func (s *UtilsTestSuite) Test_isSemverLike_basicChecks_isSemver() {
	isSemvers := []string{
		"0.0.0",
		"1.0.0",
		"10.0.0",
		"0.0.0",
		"0.1.0",
		"0.10.0",
		"0.0.0",
		"0.0.1",
		"0.0.10",
		"0.0.0-alpha",
		"0.0.0-alpha.1",
		"0.0.0-beta",
		"0.0.0-beta.1",
		"0.0.0-rc",
		"0.0.0-rc.42",
	}
	for _, isSemver := range isSemvers {
		assert.True(s.T(), isSemverLike(isSemver))
	}
}

func (s *UtilsTestSuite) Test_isSemverLike_basicChecks_notSemver() {
	notSemvers := []string{
		"00.0.0",
		"0.00.0",
		"0.0.00",
		"a.0.0",
		"1.b.0",
		"10.0.c",
		"a.0.0",
		"0.1.b",
		"0.10.c",
		"!.0.0",
		"0.X.1",
		"0.*.10",
	}
	for _, notSemver := range notSemvers {
		assert.False(s.T(), isSemverLike(notSemver))
	}
}

func (s *UtilsTestSuite) Test_isSemverLike_basicChecks_isSemverWithLabel() {
	isSemvers := []string{
		"0.0.0-label",
		"0.0.0-labelWithCaps",
		"0.0.0-label.with.dots",
		"0.0.0-label-with-dashes",
		"0.0.0-label_with_underscores",
		"0.0.0-label.with-all_possible-delims",
	}
	for _, isSemver := range isSemvers {
		assert.True(s.T(), isSemverLike(isSemver))
	}
}

func (s *UtilsTestSuite) Test_isSemverLike_basicChecks_notSemverWithLabel() {
	isSemvers := []string{
		"0.0.0label",
		"0.0.00000",
	}
	for _, isSemver := range isSemvers {
		assert.False(s.T(), isSemverLike(isSemver))
	}
}

func (s *UtilsTestSuite) Test_removeEmptyStringsFromStringSlice() {
	testSlice := []string{"", "1", "", "2", "3", ""}
	outputSlice := removeEmptyStringsFromStringSlice(testSlice)
	assert.Len(s.T(), outputSlice, 3)
	assert.Equal(s.T(), "3", outputSlice[2])
}

func (s *UtilsTestSuite) Test_toSemver() {
	testCases := map[string]ISemver{
		"1.0.0":           New(1, 0, 0, ""),
		"1.0.0-alpha.5":   New(1, 0, 0, "alpha.5"),
		"1.1.0":           New(1, 1, 0, ""),
		"1.1.1":           New(1, 1, 1, ""),
		"1.1.10":          New(1, 1, 10, ""),
		"1.10.10":         New(1, 10, 10, ""),
		"10.10.10":        New(10, 10, 10, ""),
		"10.10.10-beta.3": New(10, 10, 10, "beta.3"),
		"10.10.10-beta":   New(10, 10, 10, "beta"),
		"10.11.0":         New(10, 11, 0, ""),
	}
	for version, semver := range testCases {
		assert.Equal(s.T(), semver, toSemver(version))
	}
}

func (s *UtilsTestSuite) Test_trimAndNormalise() {
	testStrings := []string{
		"from windows\r\n",
		"from unix\n",
	}
	var outputStrings []string
	for _, testString := range testStrings {
		outputStrings = append(outputStrings, trimAndNormalise(testString))
	}
	assert.Equal(
		s.T(),
		[]string{
			"from windows",
			"from unix",
		},
		outputStrings,
	)

}
