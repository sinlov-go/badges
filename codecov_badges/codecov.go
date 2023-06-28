package codecov_badges

import "fmt"

const CodecovUrl = "https://codecov.io"

// GithubMarkdown
//
//	return markdown for codecov from github
//
// See https://docs.codecov.com/docs/status-badges
func GithubMarkdown(user, repo, branch string) string {
	return fmt.Sprintf(
		"[![codecov](%s/gh/%s/%s/branch/%s/graph/badge.svg)](%s/gh/%s/%s)",
		CodecovUrl, user, repo, branch, CodecovUrl, user, repo)
}
