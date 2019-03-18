package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SemverUtilsTestSuite struct {
	suite.Suite
}

func TestSemverUtils(t *testing.T) {
	suite.Run(t, new(SemverUtilsTestSuite))
}

func (s *SemverUtilsTestSuite) TestNew() {
	semver := New(1, 2, 3, "new.4")
	assert.Equal(s.T(), 1, semver.GetMajorInt())
	assert.Equal(s.T(), 2, semver.GetMinorInt())
	assert.Equal(s.T(), 3, semver.GetPatchInt())
	assert.Equal(s.T(), "new", semver.GetLabelString())
	assert.Equal(s.T(), 4, semver.GetLabelInt())
}

func (s *SemverUtilsTestSuite) TestNewFrom() {
	semver, err := NewFrom(func() (int, int, int, string, string, error) {
		return 5, 6, 7, "newFrom.8", "", nil
	})
	if err != nil {
		panic(err)
	}
	assert.Equal(s.T(), 5, semver.GetMajorInt())
	assert.Equal(s.T(), 6, semver.GetMinorInt())
	assert.Equal(s.T(), 7, semver.GetPatchInt())
	assert.Equal(s.T(), "newFrom", semver.GetLabelString())
	assert.Equal(s.T(), 8, semver.GetLabelInt())
}

func (s *SemverUtilsTestSuite) TestSort() {
	semvers := []ISemver{
		New(1, 0, 0, "alpha.1"),
		New(1, 0, 0, "beta"),
		New(1, 0, 0, "rc.1"),
		New(1, 0, 0, "rc.2"),
		New(1, 0, 0, "rc"),
		New(1, 1, 1, ""),
		New(1, 1, 1, "beta.1"),
		New(2, 10, 0, ""),
	}
	sortedSemvers := Sort(semvers)
	assert.Equal(s.T(), []ISemver{
		New(1, 0, 0, "alpha.1"),
		New(1, 0, 0, "beta"),
		New(1, 0, 0, "rc.1"),
		New(1, 0, 0, "rc.2"),
		New(1, 0, 0, "rc"),
		New(1, 1, 1, "beta.1"),
		New(1, 1, 1, ""),
		New(2, 10, 0, ""),
	}, sortedSemvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_major() {
	semvers := []ISemver{
		New(10, 0, 0, ""),
		New(2, 0, 0, ""),
		New(1, 0, 0, ""),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(1, 0, 0, ""),
		New(2, 0, 0, ""),
		New(10, 0, 0, ""),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_minor() {
	semvers := []ISemver{
		New(0, 10, 0, ""),
		New(0, 2, 0, ""),
		New(0, 1, 0, ""),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 1, 0, ""),
		New(0, 2, 0, ""),
		New(0, 10, 0, ""),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_patch() {
	semvers := []ISemver{
		New(0, 0, 10, ""),
		New(0, 0, 2, ""),
		New(0, 0, 1, ""),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 0, 1, ""),
		New(0, 0, 2, ""),
		New(0, 0, 10, ""),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_labelString() {
	semvers := []ISemver{
		New(0, 0, 0, "rc"),
		New(0, 0, 0, "beta"),
		New(0, 0, 0, "alpha"),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 0, 0, "alpha"),
		New(0, 0, 0, "beta"),
		New(0, 0, 0, "rc"),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_labelInt() {
	semvers := []ISemver{
		New(0, 0, 0, "rc.10"),
		New(0, 0, 0, "rc.2"),
		New(0, 0, 0, "rc.1"),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 0, 0, "rc.1"),
		New(0, 0, 0, "rc.2"),
		New(0, 0, 0, "rc.10"),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_labelIntStringMix() {
	semvers := []ISemver{
		New(0, 0, 0, "rc"),
		New(0, 0, 0, "rc.42"),
		New(0, 0, 0, "rc.1"),
		New(0, 0, 0, "beta.44"),
		New(0, 0, 0, "alpha"),
		New(0, 0, 0, "alpha.43"),
		New(0, 0, 0, "beta"),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 0, 0, "alpha.43"),
		New(0, 0, 0, "alpha"),
		New(0, 0, 0, "beta.44"),
		New(0, 0, 0, "beta"),
		New(0, 0, 0, "rc.1"),
		New(0, 0, 0, "rc.42"),
		New(0, 0, 0, "rc"),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_labelIntStringMix_multiSegment() {
	semvers := []ISemver{
		New(0, 0, 0, "a"),
		New(0, 0, 0, "a.b.c.d"),
		New(0, 0, 0, "a.b"),
		New(0, 0, 0, ""),
		New(0, 0, 0, "a.b.c"),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 0, 0, "a.b.c.d"),
		New(0, 0, 0, "a.b.c"),
		New(0, 0, 0, "a.b"),
		New(0, 0, 0, "a"),
		New(0, 0, 0, ""),
	}, semvers)
}

func (s *SemverUtilsTestSuite) TestBySemver_withDuplicates() {
	semvers := []ISemver{
		New(0, 0, 10, ""),
		New(0, 0, 1, ""),
		New(0, 0, 1, ""),
		New(10, 0, 1, ""),
		New(1, 10, 1, ""),
		New(1, 1, 10, ""),
		New(1, 1, 10, ""),
		New(1, 1, 1, ""),
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		New(0, 0, 1, ""),
		New(0, 0, 1, ""),
		New(0, 0, 10, ""),
		New(1, 1, 1, ""),
		New(1, 1, 10, ""),
		New(1, 1, 10, ""),
		New(1, 10, 1, ""),
		New(10, 0, 1, ""),
	}, semvers)
}
