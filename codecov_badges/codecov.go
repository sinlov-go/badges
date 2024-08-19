package codecov_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

const CodecovUrl = "https://codecov.io"

// Github
// see https://docs.codecov.com/docs/status-badges
func Github(user, repo, branch string) string {
	return fmt.Sprintf(
		"%s/gh/%s/%s/branch/%s/graph/badge.svg",
		CodecovUrl, user, repo, branch)
}

// GithubMarkdown
//
//	return markdown for codecov from github
//
// See https://docs.codecov.com/docs/status-badges
func GithubMarkdown(user, repo, branch string) string {
	return fmt.Sprintf(
		"[![codecov](%s)](%s/gh/%s/%s)",
		Github(user, repo, branch), CodecovUrl, user, repo)
}

// GithubHtmlMarkdown
// return html markdown for codecov from github
//
//	See https://docs.codecov.com/docs/status-badges
//
// size: badges.MarkdownImgSizes
func GithubHtmlMarkdown(user, repo, branch, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, Github(user, repo, branch), badges.MarkdownImgAlign, size, "codecov", "codecov")
	return fmt.Sprintf(
		"[%s](%s/gh/%s/%s)",
		htmlContent, CodecovUrl, user, repo)
}
