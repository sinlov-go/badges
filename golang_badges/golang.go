package golang_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// GithubGoModVersion
// See https://shields.io/badges/git-hub-go-mod-go-version-subdirectory-of-monorepo
func GithubGoModVersion(user, repo string) string {
	return fmt.Sprintf(
		"%s/github/go-mod/go-version/%s/%s?label=go.mod",
		badges.ShieldsUrl, user, repo)
}

// GithubGoModVersionMarkdown
// See https://shields.io/badges/git-hub-go-mod-go-version-subdirectory-of-monorepo
func GithubGoModVersionMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![go mod version](%s)](https://github.com/%s/%s)",
		GithubGoModVersion(user, repo), user, repo)
}

// GithubGoDoc
// See: https://pkg.go.dev
func GithubGoDoc(user, repo string) string {
	return fmt.Sprintf(
		"https://godoc.org/github.com/%s/%s?status.png",
		user, repo)
}

// GithubGoDocMarkdown
// See: https://pkg.go.dev
func GithubGoDocMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GoDoc](%s)](https://godoc.org/github.com/%s/%s)",
		GithubGoDoc(user, repo), user, repo)
}

// GithubGoReportCard
// See: https://goreportcard.com
func GithubGoReportCard(user, repo string) string {
	return fmt.Sprintf(
		"https://goreportcard.com/badge/github.com/%s/%s",
		user, repo)
}

// GithubGoReportCardMarkdown
// See: https://goreportcard.com
func GithubGoReportCardMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![goreportcard](%s)](https://goreportcard.com/report/github.com/%s/%s)",
		GithubGoReportCard(user, repo), user, repo)
}
