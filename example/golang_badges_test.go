package example

import (
	"github.com/sinlov-go/badges/golang_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGolang_badges(t *testing.T) {

	gitUser := "sinlov-go"
	gitRepo := "badges"

	assert.Equal(t, "[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/badges?label=go.mod)](https://github.com/sinlov-go/badges)",
		golang_badges.GithubGoModVersionMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, "[![GoDoc](https://godoc.org/github.com/sinlov-go/badges?status.png)](https://godoc.org/github.com/sinlov-go/badges)",
		golang_badges.GithubGoDocMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, "[![goreportcard](https://goreportcard.com/badge/github.com/sinlov-go/badges)](https://goreportcard.com/report/github.com/sinlov-go/badges)",
		golang_badges.GithubGoReportCardMarkdown(gitUser, gitRepo),
	)
}
