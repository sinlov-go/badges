package example

import (
	"github.com/sinlov-go/badges/codecov_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodecov_badges(t *testing.T) {
	gitUser := "sinlov-go"
	gitRepo := "badges"

	assert.Equal(t, "[![codecov](https://codecov.io/gh/sinlov-go/badges/branch/master/graph/badge.svg)](https://codecov.io/gh/sinlov-go/badges)",
		codecov_badges.GithubMarkdown(gitUser, gitRepo, "master"),
	)
}
