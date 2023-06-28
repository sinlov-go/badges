package shields_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// GithubLicenseMarkdown
// See: https://shields.io/badges/git-hub
func GithubLicenseMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub license](%s/github/license/%s/%s)](https://github.com/%s/%s)",
		badges.ShieldsUrl, user, repo, user, repo)
}

// GithubLatestSemVerTagMarkdown
//
//	return markdown for github tag latest
//
// See: https://shields.io/badges/git-hub-tag-latest-sem-ver-pre-release
func GithubLatestSemVerTagMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub latest SemVer tag)](%s/github/v/tag/%s/%s)](https://github.com/%s/%s/tags)",
		badges.ShieldsUrl, user, repo, user, repo)
}

// GithubReleaseMarkdown
// See: https://shields.io/badges/git-hub-release-release-name-instead-of-tag-name
func GithubReleaseMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub release)](%s/github/v/release/%s/%s)](https://github.com/%s/%s/releases)",
		badges.ShieldsUrl, user, repo, user, repo)
}

// GithubContributorsMarkdown
// See: https://shields.io/badges/git-hub-contributors
func GithubContributorsMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub contributors)](%s/github/contributors-anon/%s/%s)](https://github.com/%s/%s/graphs/contributors)",
		badges.ShieldsUrl, user, repo, user, repo,
	)
}
