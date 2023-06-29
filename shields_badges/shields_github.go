package shields_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// GithubLicense
// See: https://shields.io/badges/git-hub
func GithubLicense(user, repo string) string {
	return fmt.Sprintf(
		"%s/github/license/%s/%s",
		badges.ShieldsUrl, user, repo)
}

// GithubLicenseMarkdown
// See: https://shields.io/badges/git-hub
func GithubLicenseMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub license](%s)](https://github.com/%s/%s)",
		GithubLicense(user, repo), user, repo)
}

// GithubLatestSemVerTag
// See: https://shields.io/badges/git-hub-tag-latest-sem-ver-pre-release
func GithubLatestSemVerTag(user, repo string) string {
	return fmt.Sprintf(
		"%s/github/v/tag/%s/%s",
		badges.ShieldsUrl, user, repo)
}

// GithubLatestSemVerTagMarkdown
//
//	return markdown for github tag latest
//
// See: https://shields.io/badges/git-hub-tag-latest-sem-ver-pre-release
func GithubLatestSemVerTagMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub latest SemVer tag)](%s)](https://github.com/%s/%s/tags)",
		GithubLatestSemVerTag(user, repo), user, repo)
}

// GithubRelease
// See: https://shields.io/badges/git-hub-release-release-name-instead-of-tag-name
func GithubRelease(user, repo string) string {
	return fmt.Sprintf(
		"%s/github/v/release/%s/%s",
		badges.ShieldsUrl, user, repo)
}

// GithubReleaseMarkdown
// See: https://shields.io/badges/git-hub-release-release-name-instead-of-tag-name
func GithubReleaseMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub release)](%s)](https://github.com/%s/%s/releases)",
		GithubRelease(user, repo), user, repo)
}

// GithubContributors
// See: https://shields.io/badges/git-hub-contributors
func GithubContributors(user, repo string) string {
	return fmt.Sprintf(
		"%s/github/contributors-anon/%s/%s",
		badges.ShieldsUrl, user, repo,
	)
}

// GithubContributorsMarkdown
// See: https://shields.io/badges/git-hub-contributors
func GithubContributorsMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![GitHub contributors)](%s)](https://github.com/%s/%s/graphs/contributors)",
		GithubContributors(user, repo), user, repo,
	)
}
