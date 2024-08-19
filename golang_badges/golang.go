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

// GithubGoModVersionHtmlMarkdown
//
//	See https://shields.io/badges/git-hub-go-mod-go-version-subdirectory-of-monorepo
//
// size: badges.MarkdownImgSizes
func GithubGoModVersionHtmlMarkdown(user, repo string, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, GithubGoModVersion(user, repo), badges.MarkdownImgAlign, size, "go mod version", "go mod version")
	return fmt.Sprintf(
		"[%s](https://github.com/%s/%s)",
		htmlContent, user, repo)
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

// GithubGoDocHtmlMarkdown
//
//	See: https://pkg.go.dev
//
// size: badges.MarkdownImgSizes
func GithubGoDocHtmlMarkdown(user, repo string, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, GithubGoDoc(user, repo), badges.MarkdownImgAlign, size, "GoDoc", "GoDoc")
	return fmt.Sprintf(
		"[%s](https://godoc.org/github.com/%s/%s)",
		htmlContent, user, repo)
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

// GithubGoReportCardHtmlMarkdown
//
//	see: https://goreportcard.com
//
// size: badges.MarkdownImgSizes
func GithubGoReportCardHtmlMarkdown(user, repo string, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, GithubGoReportCard(user, repo), badges.MarkdownImgAlign, size, "goreportcard", "goreportcard")
	return fmt.Sprintf(
		"[%s](https://goreportcard.com/report/github.com/%s/%s)",
		htmlContent, user, repo)
}
