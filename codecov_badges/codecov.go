package codecov_badges

import "fmt"

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
