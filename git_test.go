package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GitLoaderTestSuite struct {
	suite.Suite
}

func TestGitLoader(t *testing.T) {
	suite.Run(t, new(GitLoaderTestSuite))
}
