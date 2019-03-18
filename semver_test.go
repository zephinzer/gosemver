package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SemverTestSuite struct {
	suite.Suite
	semver ISemver
}

func TestSemver(t *testing.T) {
	suite.Run(t, new(SemverTestSuite))
}

func (s *SemverTestSuite) SetupTest() {
	s.semver = &Semver{1, 2, 3, "label", ""}
}

func (s *SemverTestSuite) TestBumpMajor() {
	s.semver.BumpMajor()
	assert.Equal(s.T(), 2, s.semver.GetMajorInt())
	assert.Equal(s.T(), 0, s.semver.GetMinorInt())
	assert.Equal(s.T(), 0, s.semver.GetPatchInt())
	assert.Len(s.T(), s.semver.GetLabel(), 0)
}

func (s *SemverTestSuite) TestBumpMinor() {
	s.semver.BumpMinor()
	assert.Equal(s.T(), 1, s.semver.GetMajorInt())
	assert.Equal(s.T(), 3, s.semver.GetMinorInt())
	assert.Equal(s.T(), 0, s.semver.GetPatchInt())
	assert.Len(s.T(), s.semver.GetLabel(), 0)
}

func (s *SemverTestSuite) TestBumpPatch() {
	s.semver.BumpPatch()
	assert.Equal(s.T(), 1, s.semver.GetMajorInt())
	assert.Equal(s.T(), 2, s.semver.GetMinorInt())
	assert.Equal(s.T(), 4, s.semver.GetPatchInt())
	assert.Len(s.T(), s.semver.GetLabel(), 0)
}

func (s *SemverTestSuite) TestBumpLabel_withExistingLabelWithoutVersion() {
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.0", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestBumpLabel_withExistingLabelWithExistingVersion() {
	s.semver = New(1, 2, 3, "label.0")
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.1", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestBumpLabel_withExistingLabelWithMultipleSegments() {
	s.semver = New(1, 2, 3, "label.0.0")
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.0.1", s.semver.GetLabel())
	s.semver = New(1, 2, 3, "label.a.0")
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.a.1", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestBumpLabel_withNoLabel() {
	s.semver.BumpLabel("nonlabel")
	assert.Equal(s.T(), "nonlabel.0", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestGetMajorInt() {
	assert.Equal(s.T(), 1, s.semver.GetMajorInt())
}

func (s *SemverTestSuite) TestGetMinorInt() {
	assert.Equal(s.T(), 2, s.semver.GetMinorInt())
}

func (s *SemverTestSuite) TestGetPatchInt() {
	assert.Equal(s.T(), 3, s.semver.GetPatchInt())
}

func (s *SemverTestSuite) TestGetLabelInt() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 1, 2, 3, "label.4", "", nil
	})
	assert.Equal(s.T(), 4, s.semver.GetLabelInt())
}

func (s *SemverTestSuite) TestGetLabelString_oneDelimiter() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 1, 2, 3, "label.4", "", nil
	})
	assert.Equal(s.T(), "label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabelString_noNumber() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 1, 2, 3, "label", "", nil
	})
	assert.Equal(s.T(), "label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabelString_multiDelimiter() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 1, 2, 3, "label.another.label.4", "", nil
	})
	assert.Equal(s.T(), "label.another.label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabelString_multiDelimiter_noNumber() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 1, 2, 3, "label.another.label", "", nil
	})
	assert.Equal(s.T(), "label.another.label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabel() {
	assert.Equal(s.T(), "label", s.semver.GetLabel())
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.0", s.semver.GetLabel())
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.1", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestGetPrefix() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 1, 2, 3, "label", "prefix", nil
	})
	assert.Equal(s.T(), "prefix", s.semver.GetPrefix())
}

func (s *SemverTestSuite) TestLoad() {
	s.semver.Load(func() (int, int, int, string, string, error) {
		return 4, 5, 6, "label.42", "", nil
	})
	assert.Equal(s.T(), 4, s.semver.GetMajorInt())
	assert.Equal(s.T(), 5, s.semver.GetMinorInt())
	assert.Equal(s.T(), 6, s.semver.GetPatchInt())
	assert.Equal(s.T(), 42, s.semver.GetLabelInt())
	assert.Equal(s.T(), "label.42", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestString() {
	s.semver = New(7, 8, 9, "label.42")
	assert.Equal(s.T(), "7.8.9-label.42", s.semver.String())
	s.semver.BumpMajor()
	assert.Equal(s.T(), "8.0.0", s.semver.String())
	s.semver.BumpMinor()
	assert.Equal(s.T(), "8.1.0", s.semver.String())
	s.semver.BumpPatch()
	assert.Equal(s.T(), "8.1.1", s.semver.String())
	s.semver.BumpLabel("TestString")
	assert.Equal(s.T(), "8.1.1-TestString.0", s.semver.String())
	s.semver.BumpLabel("TestString")
	assert.Equal(s.T(), "8.1.1-TestString.1", s.semver.String())
}
