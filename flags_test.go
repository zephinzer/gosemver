package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type CLIFlagsTestSuite struct {
	suite.Suite
}

func TestCLIFlags(t *testing.T) {
	suite.Run(t, new(CLIFlagsTestSuite))
}

func (s *CLIFlagsTestSuite) Test_flags() {
	allFlags := flags(flagUse, flagMode)
	assert.Len(s.T(), allFlags, 2)
	assert.Equal(s.T(), "use, u", allFlags[0].(cli.StringFlag).Name)
	assert.Equal(s.T(), "mode, m", allFlags[1].(cli.StringFlag).Name)
}

func (s *CLIFlagsTestSuite) Test_flagUse() {
	flag := cli.StringFlag(flagUse().(cli.StringFlag))
	assert.NotNil(s.T(), flag.Usage)
	assert.Equal(s.T(), "use, u", flag.Name)
	assert.Equal(s.T(), "git", flag.Value)
	assert.Equal(s.T(), "USE", flag.EnvVar)
}

func (s *CLIFlagsTestSuite) Test_flagMode() {
	flag := cli.StringFlag(flagMode().(cli.StringFlag))
	assert.NotNil(s.T(), flag.Usage)
	assert.Equal(s.T(), "mode, m", flag.Name)
	assert.Equal(s.T(), "latest", flag.Value)
	assert.Equal(s.T(), "MODE", flag.EnvVar)
}

func (s *CLIFlagsTestSuite) Test_flagPrefix() {
	flag := cli.StringFlag(flagPrefix().(cli.StringFlag))
	assert.NotNil(s.T(), flag.Usage)
	assert.Equal(s.T(), "prefix, p", flag.Name)
	assert.Equal(s.T(), "", flag.Value)
	assert.Equal(s.T(), "PREFIX", flag.EnvVar)
}

func (s *CLIFlagsTestSuite) Test_flagYes() {
	flag := cli.BoolFlag(flagYes().(cli.BoolFlag))
	assert.NotNil(s.T(), flag.Usage)
	assert.Equal(s.T(), "yes, y", flag.Name)
	assert.Equal(s.T(), "YES", flag.EnvVar)
}
