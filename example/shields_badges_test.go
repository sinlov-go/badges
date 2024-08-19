package example

import (
	"github.com/sinlov-go/badges"
	"github.com/sinlov-go/badges/shields_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticBadge(t *testing.T) {
	label := "rust"
	message := "1.43.1"

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.43.1-blue",
		shields_badges.StaticBadgeBlue(label, message),
	)

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.52.1--rc.1-blue",
		shields_badges.StaticBadgeBlue(label, "1.52.1-rc.1"),
	)

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.52.1_rc.1-blue",
		shields_badges.StaticBadgeBlue(label, "1.52.1 rc.1"),
	)

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.52.1__rc.1-blue",
		shields_badges.StaticBadgeBlue(label, "1.52.1_rc.1"),
	)

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.43.1-green",
		shields_badges.StaticBadgeGreen(label, message),
	)
	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.43.1-red",
		shields_badges.StaticBadgeRed(label, message),
	)
	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.43.1-yellow",
		shields_badges.StaticBadgeYellow(label, message),
	)

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.43.1-orange",
		shields_badges.StaticBadgeOrange(label, message),
	)

	assert.Equal(t,
		"https://img.shields.io/badge/rust-1.43.1-purple",
		shields_badges.StaticBadgePurple(label, message),
	)
}

func TestShields_badges(t *testing.T) {
	gitUser := "sinlov-go"
	gitRepo := "badges"

	assert.Equal(t, "https://img.shields.io/github/license/sinlov-go/badges",
		shields_badges.GithubLicense(gitUser, gitRepo),
	)

	assert.Equal(t,
		"[![GitHub license](https://img.shields.io/github/license/sinlov-go/badges)](https://github.com/sinlov-go/badges)",
		shields_badges.GithubLicenseMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/github/license/sinlov-go/badges" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="GitHub license" title="GitHub license">](https://github.com/sinlov-go/badges)`,
		shields_badges.GithubLicenseHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/github/v/tag/sinlov-go/badges",
		shields_badges.GithubLatestSemVerTag(gitUser, gitRepo),
	)

	assert.Equal(t,
		"[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov-go/badges)](https://github.com/sinlov-go/badges/tags)",
		shields_badges.GithubLatestSemVerTagMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/github/v/tag/sinlov-go/badges" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="GitHub latest SemVer tag" title="GitHub latest SemVer tag">](https://github.com/sinlov-go/badges/tags)`,
		shields_badges.GithubLatestSemVerTagHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/github/v/release/sinlov-go/badges",
		shields_badges.GithubRelease(gitUser, gitRepo),
	)

	assert.Equal(t,
		"[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/badges)](https://github.com/sinlov-go/badges/releases)",
		shields_badges.GithubReleaseMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/github/v/release/sinlov-go/badges" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="GitHub release" title="GitHub release">](https://github.com/sinlov-go/badges/releases)`,
		shields_badges.GithubReleaseHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/github/contributors-anon/sinlov-go/badges",
		shields_badges.GithubContributors(gitUser, gitRepo),
	)

	assert.Equal(t,
		"[![GitHub contributors)](https://img.shields.io/github/contributors-anon/sinlov-go/badges)](https://github.com/sinlov-go/badges/graphs/contributors)",
		shields_badges.GithubContributorsMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/github/contributors-anon/sinlov-go/badges" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="GitHub contributors" title="GitHub contributors">](https://github.com/sinlov-go/badges/graphs/contributors)`,
		shields_badges.GithubContributorsHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)
}

func TestShieldsDocker(t *testing.T) {

	dockerUser := "sinlov"
	dockerRepo := "docker-rust-buster"

	assert.Equal(t, "https://img.shields.io/docker/v/sinlov/docker-rust-buster?sort=semver",
		shields_badges.DockerHubImageVersionSemver(dockerUser, dockerRepo),
	)

	assert.Equal(t,
		"[![docker hub version semver](https://img.shields.io/docker/v/sinlov/docker-rust-buster?sort=semver)](https://hub.docker.com/r/sinlov/docker-rust-buster/tags?page=1&ordering=last_updated)",
		shields_badges.DockerHubImageVersionSemverMarkdown(dockerUser, dockerRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/docker/v/sinlov/docker-rust-buster?sort=semver" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="docker hub version semver" title="docker hub version semver">](https://hub.docker.com/r/sinlov/docker-rust-buster/tags?page=1&ordering=last_updated)`,
		shields_badges.DockerHubImageVersionSemverHtmlMarkdown(dockerUser, dockerRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/docker/image-size/sinlov/docker-rust-buster",
		shields_badges.DockerHubImageSize(dockerUser, dockerRepo),
	)

	assert.Equal(t,
		"[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/docker-rust-buster)](https://hub.docker.com/r/sinlov/docker-rust-buster)",
		shields_badges.DockerHubImageSizeMarkdown(dockerUser, dockerRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/docker/image-size/sinlov/docker-rust-buster" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="docker hub image size" title="docker hub image size">](https://hub.docker.com/r/sinlov/docker-rust-buster)`,
		shields_badges.DockerHubImageSizeHtmlMarkdown(dockerUser, dockerRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/docker/pulls/sinlov/docker-rust-buster",
		shields_badges.DockerHubImagePull(dockerUser, dockerRepo),
	)

	assert.Equal(t,
		"[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/docker-rust-buster)](https://hub.docker.com/r/sinlov/docker-rust-buster/tags?page=1&ordering=last_updated)",
		shields_badges.DockerHubImagePullMarkdown(dockerUser, dockerRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/docker/pulls/sinlov/docker-rust-buster" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="docker hub image pulls" title="docker hub image pulls">](https://hub.docker.com/r/sinlov/docker-rust-buster/tags?page=1&ordering=last_updated)`,
		shields_badges.DockerHubImagePullHtmlMarkdown(dockerUser, dockerRepo, badges.MarkdownImgSizes),
	)
}
