package shields_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// DockerHubImageVersionSemverMarkdown
//
//	user
//	repo
//
// See: https://shields.io/badges/docker-image-version-latest-semver
func DockerHubImageVersionSemverMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub version semver](%s/docker/v/%s/%s?sort=semver)](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		badges.ShieldsUrl, user, repo, user, repo)
}

// DockerHubImageSizeMarkdown
//
//	user
//	repo
//
// See: https://shields.io/badges/docker-image-size-with-architecture-latest-by-date-latest-semver
func DockerHubImageSizeMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub image size](%s/docker/image-size/%s/%s)](https://hub.docker.com/r/%s/%s)",
		badges.ShieldsUrl, user, repo, user, repo)
}

// DockerHubImagePullMarkdown
//
//	user
//	repo
//
// See: https://shields.io/badges/docker-pulls
func DockerHubImagePullMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![docker hub image size](%s/docker/pulls/%s/%s)](https://hub.docker.com/r/%s/%s/tags?page=1&ordering=last_updated)",
		badges.ShieldsUrl, user, repo, user, repo)
}
