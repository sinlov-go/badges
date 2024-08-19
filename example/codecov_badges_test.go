package example

import (
	"github.com/sinlov-go/badges"
	"github.com/sinlov-go/badges/codecov_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodecov_badges(t *testing.T) {
	gitUser := "sinlov-go"
	gitRepo := "badges"

	assert.Equal(t, "https://codecov.io/gh/sinlov-go/badges/branch/main/graph/badge.svg",
		codecov_badges.Github(gitUser, gitRepo, "main"),
	)

	assert.Equal(t, "[![codecov](https://codecov.io/gh/sinlov-go/badges/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/badges)",
		codecov_badges.GithubMarkdown(gitUser, gitRepo, "main"),
	)

	assert.Equal(t, `[<img src="https://codecov.io/gh/sinlov-go/badges/branch/main/graph/badge.svg" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="codecov" title="codecov">](https://codecov.io/gh/sinlov-go/badges)`,
		codecov_badges.GithubHtmlMarkdown(gitUser, gitRepo, "main", badges.MarkdownImgSizes))
}
