package main

import (
	"sort"
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
	s.semver = &Semver{1, 2, 3, "label"}
}

func (s *SemverTestSuite) TestBySemver_major() {
	semvers := []ISemver{
		&Semver{10, 0, 0, ""},
		&Semver{2, 0, 0, ""},
		&Semver{1, 0, 0, ""},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{1, 0, 0, ""},
		&Semver{2, 0, 0, ""},
		&Semver{10, 0, 0, ""},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_minor() {
	semvers := []ISemver{
		&Semver{0, 10, 0, ""},
		&Semver{0, 2, 0, ""},
		&Semver{0, 1, 0, ""},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 1, 0, ""},
		&Semver{0, 2, 0, ""},
		&Semver{0, 10, 0, ""},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_patch() {
	semvers := []ISemver{
		&Semver{0, 0, 10, ""},
		&Semver{0, 0, 2, ""},
		&Semver{0, 0, 1, ""},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 0, 1, ""},
		&Semver{0, 0, 2, ""},
		&Semver{0, 0, 10, ""},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_labelString() {
	semvers := []ISemver{
		&Semver{0, 0, 0, "rc"},
		&Semver{0, 0, 0, "beta"},
		&Semver{0, 0, 0, "alpha"},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 0, 0, "alpha"},
		&Semver{0, 0, 0, "beta"},
		&Semver{0, 0, 0, "rc"},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_labelInt() {
	semvers := []ISemver{
		&Semver{0, 0, 0, "rc.10"},
		&Semver{0, 0, 0, "rc.2"},
		&Semver{0, 0, 0, "rc.1"},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 0, 0, "rc.1"},
		&Semver{0, 0, 0, "rc.2"},
		&Semver{0, 0, 0, "rc.10"},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_labelIntStringMix() {
	semvers := []ISemver{
		&Semver{0, 0, 0, "rc"},
		&Semver{0, 0, 0, "rc.42"},
		&Semver{0, 0, 0, "rc.1"},
		&Semver{0, 0, 0, "beta.44"},
		&Semver{0, 0, 0, "alpha"},
		&Semver{0, 0, 0, "alpha.43"},
		&Semver{0, 0, 0, "beta"},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 0, 0, "alpha.43"},
		&Semver{0, 0, 0, "alpha"},
		&Semver{0, 0, 0, "beta.44"},
		&Semver{0, 0, 0, "beta"},
		&Semver{0, 0, 0, "rc.1"},
		&Semver{0, 0, 0, "rc.42"},
		&Semver{0, 0, 0, "rc"},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_labelIntStringMix_multiSegment() {
	semvers := []ISemver{
		&Semver{0, 0, 0, "a"},
		&Semver{0, 0, 0, "a.b.c.d"},
		&Semver{0, 0, 0, "a.b"},
		&Semver{0, 0, 0, ""},
		&Semver{0, 0, 0, "a.b.c"},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 0, 0, "a.b.c.d"},
		&Semver{0, 0, 0, "a.b.c"},
		&Semver{0, 0, 0, "a.b"},
		&Semver{0, 0, 0, "a"},
		&Semver{0, 0, 0, ""},
	}, semvers)
}

func (s *SemverTestSuite) TestBySemver_withDuplicates() {
	semvers := []ISemver{
		&Semver{0, 0, 10, ""},
		&Semver{0, 0, 1, ""},
		&Semver{0, 0, 1, ""},
		&Semver{10, 0, 1, ""},
		&Semver{1, 10, 1, ""},
		&Semver{1, 1, 10, ""},
		&Semver{1, 1, 10, ""},
		&Semver{1, 1, 1, ""},
	}
	sort.Stable(BySemver(semvers))
	assert.Equal(s.T(), []ISemver{
		&Semver{0, 0, 1, ""},
		&Semver{0, 0, 1, ""},
		&Semver{0, 0, 10, ""},
		&Semver{1, 1, 1, ""},
		&Semver{1, 1, 10, ""},
		&Semver{1, 1, 10, ""},
		&Semver{1, 10, 1, ""},
		&Semver{10, 0, 1, ""},
	}, semvers)
}

func (s *SemverTestSuite) TestNew() {
	semver := New(1, 2, 3, "new.4")
	assert.Equal(s.T(), 1, semver.GetMajorInt())
	assert.Equal(s.T(), 2, semver.GetMinorInt())
	assert.Equal(s.T(), 3, semver.GetPatchInt())
	assert.Equal(s.T(), "new", semver.GetLabelString())
	assert.Equal(s.T(), 4, semver.GetLabelInt())
}

func (s *SemverTestSuite) TestNewFrom() {
	semver := NewFrom(func() (int, int, int, string, error) {
		return 5, 6, 7, "newFrom.8", nil
	})
	assert.Equal(s.T(), 5, semver.GetMajorInt())
	assert.Equal(s.T(), 6, semver.GetMinorInt())
	assert.Equal(s.T(), 7, semver.GetPatchInt())
	assert.Equal(s.T(), "newFrom", semver.GetLabelString())
	assert.Equal(s.T(), 8, semver.GetLabelInt())
}

func (s *SemverTestSuite) TestSort() {
	semvers := []ISemver{
		&Semver{1, 0, 0, "alpha.1"},
		&Semver{1, 0, 0, "beta"},
		&Semver{1, 0, 0, "rc.1"},
		&Semver{1, 0, 0, "rc.2"},
		&Semver{1, 0, 0, "rc"},
		&Semver{1, 1, 1, ""},
		&Semver{1, 1, 1, "beta.1"},
		&Semver{2, 10, 0, ""},
	}
	sortedSemvers := Sort(semvers)
	assert.Equal(s.T(), []ISemver{
		&Semver{1, 0, 0, "alpha.1"},
		&Semver{1, 0, 0, "beta"},
		&Semver{1, 0, 0, "rc.1"},
		&Semver{1, 0, 0, "rc.2"},
		&Semver{1, 0, 0, "rc"},
		&Semver{1, 1, 1, "beta.1"},
		&Semver{1, 1, 1, ""},
		&Semver{2, 10, 0, ""},
	}, sortedSemvers)
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
	s.semver = &Semver{1, 2, 3, "label.0"}
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.1", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestBumpLabel_withExistingLabelWithMultipleSegments() {
	s.semver = &Semver{1, 2, 3, "label.0.0"}
	s.semver.BumpLabel("label")
	assert.Equal(s.T(), "label.0.1", s.semver.GetLabel())
	s.semver = &Semver{1, 2, 3, "label.a.0"}
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
	s.semver.Load(func() (int, int, int, string, error) {
		return 1, 2, 3, "label.4", nil
	})
	assert.Equal(s.T(), 4, s.semver.GetLabelInt())
}

func (s *SemverTestSuite) TestGetLabelString_oneDelimiter() {
	s.semver.Load(func() (int, int, int, string, error) {
		return 1, 2, 3, "label.4", nil
	})
	assert.Equal(s.T(), "label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabelString_noNumber() {
	s.semver.Load(func() (int, int, int, string, error) {
		return 1, 2, 3, "label", nil
	})
	assert.Equal(s.T(), "label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabelString_multiDelimiter() {
	s.semver.Load(func() (int, int, int, string, error) {
		return 1, 2, 3, "label.another.label.4", nil
	})
	assert.Equal(s.T(), "label.another.label", s.semver.GetLabelString())
}

func (s *SemverTestSuite) TestGetLabelString_multiDelimiter_noNumber() {
	s.semver.Load(func() (int, int, int, string, error) {
		return 1, 2, 3, "label.another.label", nil
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

func (s *SemverTestSuite) TestLoad() {
	s.semver.Load(func() (int, int, int, string, error) {
		return 4, 5, 6, "label.42", nil
	})
	assert.Equal(s.T(), 4, s.semver.GetMajorInt())
	assert.Equal(s.T(), 5, s.semver.GetMinorInt())
	assert.Equal(s.T(), 6, s.semver.GetPatchInt())
	assert.Equal(s.T(), 42, s.semver.GetLabelInt())
	assert.Equal(s.T(), "label.42", s.semver.GetLabel())
}

func (s *SemverTestSuite) TestString() {
	s.semver = &Semver{7, 8, 9, "label.42"}
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
