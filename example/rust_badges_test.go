package example

import (
	"fmt"
	"github.com/sinlov-go/badges"
	"github.com/sinlov-go/badges/rust_badges"
	"github.com/sinlov-go/badges/shields_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithubRust(t *testing.T) {
	t.Logf("~> mock GithubRust")
	// mock GithubRust

	gitUser := "seanmonstar"
	gitRepo := "reqwest"
	rustVersion := "1.65"

	assert.Equal(t, "https://img.shields.io/crates/v/reqwest",
		rust_badges.CratesVersion(gitRepo),
	)

	assert.Equal(t, "[![crates.io version](https://img.shields.io/crates/v/reqwest)](https://crates.io/crates/reqwest)",
		rust_badges.CratesVersionMarkdown(gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/crates/v/reqwest" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="crates.io version" title="crates.io version">](https://crates.io/crates/reqwest)`,
		rust_badges.CratesVersionHtmlMarkdown(gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/crates/l/reqwest",
		rust_badges.CratesLicense(gitRepo),
	)

	assert.Equal(t, "[![crates.io license](https://img.shields.io/crates/l/reqwest)](https://crates.io/crates/reqwest)",
		rust_badges.CratesLicenseMarkdown(gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/crates/l/reqwest" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="crates.io license" title="crates.io license">](https://crates.io/crates/reqwest)`,
		rust_badges.CratesLicenseHtmlMarkdown(gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/docsrs/reqwest",
		rust_badges.DocsRs(gitRepo),
	)

	assert.Equal(t, "[![doc.rs](https://img.shields.io/docsrs/reqwest)](https://docs.rs/reqwest)",
		rust_badges.DocsRsMarkdown(gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/docsrs/reqwest" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="doc.rs" title="doc.rs">](https://docs.rs/reqwest)`,
		rust_badges.DocsRsHtmlMarkdown(gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/crates/dv/reqwest.svg",
		rust_badges.CratesDownloadLatest(gitRepo),
	)

	assert.Equal(t, "[![crates.io download latest](https://img.shields.io/crates/dv/reqwest.svg)](https://crates.io/crates/reqwest)",
		rust_badges.CratesDownloadLatestMarkdown(gitRepo),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/crates/dv/reqwest.svg" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="crates.io download latest" title="crates.io download latest">](https://crates.io/crates/reqwest)`,
		rust_badges.CratesDownloadLatestHtmlMarkdown(gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://deps.rs/crate/reqwest/latest/status.svg",
		rust_badges.DepsRsCrateLatest(gitRepo),
	)

	assert.Equal(t, "[![deps.rs dependency](https://deps.rs/crate/reqwest/latest/status.svg)](https://deps.rs/repo/github/reqwest/latest)",
		rust_badges.DepsRsCrateLatestMarkdown(gitRepo),
	)

	assert.Equal(t, `[<img src="https://deps.rs/crate/reqwest/latest/status.svg" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="deps.rs dependency" title="deps.rs dependency">](https://deps.rs/repo/github/reqwest/latest)`,
		rust_badges.DepsRsCrateLatestHtmlMarkdown(gitRepo, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://deps.rs/repo/github/seanmonstar/reqwest/status.svg",
		rust_badges.DepsRsGithub(gitUser, gitRepo),
	)

	assert.Equal(t, "[![deps.rs dependency](https://deps.rs/repo/github/seanmonstar/reqwest/status.svg)](https://deps.rs/repo/github/seanmonstar/reqwest)",
		rust_badges.DepsRsGithubMarkdown(gitUser, gitRepo),
	)

	assert.Equal(t, `[<img src="https://deps.rs/repo/github/seanmonstar/reqwest/status.svg" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="deps.rs dependency" title="deps.rs dependency">](https://deps.rs/repo/github/seanmonstar/reqwest)`,
		rust_badges.DepsRsGithubHtmlMarkdown(gitUser, gitRepo, badges.MarkdownImgSizes),
	)

	// verify GithubRust
	assert.Equal(t, "[![rust version](https://img.shields.io/badge/rust-1.65-orange)](https://github.com/seanmonstar/reqwest)",
		fmt.Sprintf("[![rust version](%s)](https://github.com/%s/%s)",
			shields_badges.StaticBadgeOrange("rust", rustVersion),
			gitUser, gitRepo,
		),
	)
}
