package rust_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// CratesVersion
// See: https://shields.io/badges/crates-io-3
func CratesVersion(repo string) string {
	return fmt.Sprintf(
		"%s/crates/v/%s", badges.ShieldsUrl, repo,
	)
}

// CratesVersionMarkdown
// See: https://shields.io/badges/crates-io-3
func CratesVersionMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![crates.io version](%s)](https://crates.io/crates/%s)",
		CratesVersion(repo), repo,
	)
}

// CratesLicense
// See: https://shields.io/badges/crates-io-1
func CratesLicense(repo string) string {
	return fmt.Sprintf(
		"%s/crates/l/%s", badges.ShieldsUrl, repo,
	)
}

// CratesLicenseMarkdown
// See: https://shields.io/badges/crates-io-1
func CratesLicenseMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![crates.io license](%s)](https://crates.io/crates/%s)",
		CratesLicense(repo), repo,
	)
}

// DocsRs
// See: https://shields.io/badges/docs-rs-1
func DocsRs(repo string) string {
	return fmt.Sprintf(
		"%s/docsrs/%s", badges.ShieldsUrl, repo,
	)
}

// DocsRsMarkdown
// See: https://shields.io/badges/docs-rs-1
func DocsRsMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![doc.rs](%s)](https://docs.rs/%s)",
		DocsRs(repo), repo,
	)
}

// CratesDownloadLatest
// See: https://shields.io/badges/crates-io-latest
func CratesDownloadLatest(repo string) string {
	return fmt.Sprintf(
		"%s/crates/dv/%s.svg", badges.ShieldsUrl, repo,
	)
}

// CratesDownloadLatestMarkdown
// See: https://shields.io/badges/crates-io-latest
func CratesDownloadLatestMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![crates.io download latest](%s)](https://crates.io/crates/%s)",
		CratesDownloadLatest(repo), repo,
	)
}

// DepsRsCrateLatest
// See: https://deps.rs/
func DepsRsCrateLatest(repo string) string {
	return fmt.Sprintf(
		"https://deps.rs/crate/%s/latest/status.svg", repo,
	)
}

// DepsRsCrateLatestMarkdown
// See: https://deps.rs/
func DepsRsCrateLatestMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![deps.rs dependency](%s)](https://deps.rs/repo/github/%s/latest)",
		DepsRsCrateLatest(repo), repo)
}

// DepsRsGithub
// See: https://deps.rs/
func DepsRsGithub(user, repo string) string {
	return fmt.Sprintf(
		"https://deps.rs/repo/github/%s/%s/status.svg", user, repo,
	)
}

// DepsRsGithubMarkdown
// see https://deps.rs/
func DepsRsGithubMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![deps.rs dependency](%s)](https://deps.rs/repo/github/%s/%s)",
		DepsRsGithub(user, repo), user, repo)
}
