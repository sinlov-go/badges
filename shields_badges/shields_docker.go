package shields_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// DockerHubImageVersionSemver
//
//	See: https://shields.io/badges/docker-image-version-latest-semver
//
// user: docker hub user
// repo: docker hub repo
func DockerHubImageVersionSemver(user, repo string) string {
	return fmt.Sprintf(
		"%s/docker/v/%s/%s?sort=semver",
		badges.ShieldsUrl, user, repo)
}

// DockerHubImageVersionSemverMarkdown
//
//	See: https://shields.io/badges/docker-image-version-latest-semver
//
// user: docker hub user
// repo: docker hub repo
func DockerHubImageVersionSemverMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub version semver](%s)](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		DockerHubImageVersionSemver(user, repo), user, repo)
}

// DockerHubImageVersionSemverHtmlMarkdown
//
//	see DockerHubImageVersionSemverHtmlMarkdown
//
// user: docker hub use
// repo: docker hub repo
// size: badges.MarkdownImgSizes
func DockerHubImageVersionSemverHtmlMarkdown(user, repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DockerHubImageVersionSemver(user, repo), badges.MarkdownImgAlign, size, "docker hub version semver", "docker hub version semver")
	return fmt.Sprintf(
		"[%s](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		htmlContent, user, repo)
}

// DockerHubImageSize
// See: https://shields.io/badges/docker-image-size-with-architecture-latest-by-date-latest-semver
func DockerHubImageSize(user, repo string) string {
	return fmt.Sprintf(
		"%s/docker/image-size/%s/%s",
		badges.ShieldsUrl, user, repo)
}

// DockerHubImageSizeMarkdown
//
//	See: https://shields.io/badges/docker-image-size-with-architecture-latest-by-date-latest-semver
//
// user: docker hub user
// repo: docker hub repo
func DockerHubImageSizeMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub image size](%s)](https://hub.docker.com/r/%s/%s)",
		DockerHubImageSize(user, repo), user, repo)
}

// DockerHubImageSizeHtmlMarkdown
//
//	see https://shields.io/badges/docker-image-size-with-architecture-latest-by-date-latest-semver
//
// user: docker hub user
// repo: docker hub repo
// size: badges.MarkdownImgSizes
func DockerHubImageSizeHtmlMarkdown(user, repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DockerHubImageSize(user, repo), badges.MarkdownImgAlign, size, "docker hub image size", "docker hub image size")
	return fmt.Sprintf(
		"[%s](https://hub.docker.com/r/%s/%s)",
		htmlContent, user, repo)
}

// DockerHubImagePull
//
//	See: https://shields.io/badges/docker-pulls
//
// user: docker hub user
// repo: docker hub repo
func DockerHubImagePull(user, repo string) string {
	return fmt.Sprintf(
		"%s/docker/pulls/%s/%s",
		badges.ShieldsUrl, user, repo)
}

// DockerHubImagePullMarkdown
//
//	See: https://shields.io/badges/docker-pulls
//
// repo: docker hub repo
// user: docker hub user
func DockerHubImagePullMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub image pulls](%s)](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		DockerHubImagePull(user, repo), user, repo)
}

// DockerHubImagePullHtmlMarkdown
//
//	see https://shields.io/badges/docker-pulls
//
// user: docker hub user
// repo: docker hub repo
// size: badges.MarkdownImgSizes
func DockerHubImagePullHtmlMarkdown(user, repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DockerHubImagePull(user, repo), badges.MarkdownImgAlign, size, "docker hub image pulls", "docker hub image pulls")
	return fmt.Sprintf(
		"[%s](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		htmlContent, user, repo)
}
