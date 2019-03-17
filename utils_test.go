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
