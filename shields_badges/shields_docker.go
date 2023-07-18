package shields_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// DockerHubImageVersionSemver
// See: https://shields.io/badges/docker-image-version-latest-semver
func DockerHubImageVersionSemver(user, repo string) string {
	return fmt.Sprintf(
		"%s/docker/v/%s/%s?sort=semver",
		badges.ShieldsUrl, user, repo)
}

// DockerHubImageVersionSemverMarkdown
//
//	user
//	repo
//
// See: https://shields.io/badges/docker-image-version-latest-semver
func DockerHubImageVersionSemverMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub version semver](%s)](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		DockerHubImageVersionSemver(user, repo), user, repo)
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
//	user
//	repo
//
// See: https://shields.io/badges/docker-image-size-with-architecture-latest-by-date-latest-semver
func DockerHubImageSizeMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub image size](%s)](https://hub.docker.com/r/%s/%s)",
		DockerHubImageSize(user, repo), user, repo)
}

// DockerHubImagePull
// See: https://shields.io/badges/docker-pulls
func DockerHubImagePull(user, repo string) string {
	return fmt.Sprintf(
		"%s/docker/pulls/%s/%s",
		badges.ShieldsUrl, user, repo)
}

// DockerHubImagePullMarkdown
//
//	user
//	repo
//
// See: https://shields.io/badges/docker-pulls
func DockerHubImagePullMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub image pulls](%s)](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		DockerHubImagePull(user, repo), user, repo)
}
