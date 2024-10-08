package example

import (
	"github.com/sinlov-go/badges"
	"github.com/sinlov-go/badges/golang_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGolang_badges(t *testing.T) {

	gitUser := "sinlov-go"
	gitRepo := "badges"

	assert.Equal(t, "https://img.shields.io/github/go-mod/go-version/sinlov-go/badges?label=go.mod",
		golang_badges.GithubGoModVersion(gitUser, gitRepo),
	)

	assert.Equal(t, "[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/badges?label=go.mod)](https://github.com/sinlov-go/badges)",
		golang_badges.GithubGoModVersionMarkdown(gitUser, gitRepo),
	)
	assert.Equal(t, `[<img src="https://img.shields.io/github/go-mod/go-version/sinlov-go/badges?label=go.mod" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="go mod version" title="go mod version">](https://github.com/sinlov-go/badges)`,
		golang_badges.GithubGoModVersionHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://godoc.org/github.com/sinlov-go/badges?status.png",
		golang_badges.GithubGoDoc(gitUser, gitRepo),
	)

	assert.Equal(t, "[![GoDoc](https://godoc.org/github.com/sinlov-go/badges?status.png)](https://godoc.org/github.com/sinlov-go/badges)",
		golang_badges.GithubGoDocMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://godoc.org/github.com/sinlov-go/badges?status.png" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="GoDoc" title="GoDoc">](https://godoc.org/github.com/sinlov-go/badges)`,
		golang_badges.GithubGoDocHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://goreportcard.com/badge/github.com/sinlov-go/badges",
		golang_badges.GithubGoReportCard(gitUser, gitRepo),
	)

	assert.Equal(t, "[![goreportcard](https://goreportcard.com/badge/github.com/sinlov-go/badges)](https://goreportcard.com/report/github.com/sinlov-go/badges)",
		golang_badges.GithubGoReportCardMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://goreportcard.com/badge/github.com/sinlov-go/badges" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="goreportcard" title="goreportcard">](https://goreportcard.com/report/github.com/sinlov-go/badges)`,
		golang_badges.GithubGoReportCardHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)
}
