package rust_badges

import "fmt"

// CratesVersionMarkdown
// see: https://shields.io/badges/crates-io-3
func CratesVersionMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![crates.io version](https://img.shields.io/crates/v/%s)](https://crates.io/crates/%s)",
		repo, repo,
	)
}

// CratesLicenseMarkdown
// See: https://shields.io/badges/crates-io-1
func CratesLicenseMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![crates.io license](https://img.shields.io/crates/l/%s)](https://crates.io/crates/%s)",
		repo, repo,
	)
}

// DocsRsMarkdown
// see: https://shields.io/badges/docs-rs-1
func DocsRsMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![doc.rs](https://img.shields.io/docsrs/%s)](https://docs.rs//%s)",
		repo, repo,
	)
}

// CratesDownloadLatestMarkdown
// See: https://shields.io/badges/crates-io-latest
func CratesDownloadLatestMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![crates.io download latest](https://img.shields.io/crates/dv/%s.svg)](https://crates.io/crates/%s)",
		repo, repo,
	)
}

// DepsRsCrateLatestMarkdown
// see: https://deps.rs/
func DepsRsCrateLatestMarkdown(repo string) string {
	return fmt.Sprintf(
		"[![deps.rs dependency](https://deps.rs/crate/%s/latest/status.svg)](https://deps.rs/repo/github/%s/latest)",
		repo, repo)
}

// DepsRsGithubMarkdown
// see https://deps.rs/
func DepsRsGithubMarkdown(user, repo string) string {
	return fmt.Sprintf(
		"[![deps.rs dependency](https://deps.rs/repo/github/%s/%s/status.svg)](https://deps.rs/repo/github/%s/%s)",
		user, repo, user, repo)
}
