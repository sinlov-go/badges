package golang_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// GithubGoModVersionMarkdown
// See https://shields.io/badges/git-hub-go-mod-go-version-subdirectory-of-monorepo
func GithubGoModVersionMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![go mod version](%s/github/go-mod/go-version/%s/%s?label=go.mod)](https://github.com/%s/%s)",
		badges.ShieldsUrl, user, repo, user, repo)
}

// GithubGoDocMarkdown
// See: https://pkg.go.dev
func GithubGoDocMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GoDoc](https://godoc.org/github.com/%s/%s?status.png)](https://godoc.org/github.com/%s/%s)",
		user, repo, user, repo)
}

// GithubGoReportCardMarkdown
// See: https://goreportcard.com
func GithubGoReportCardMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![goreportcard](https://goreportcard.com/badge/github.com/%s/%s)](https://goreportcard.com/report/github.com/%s/%s)",
		user, repo, user, repo)
}
