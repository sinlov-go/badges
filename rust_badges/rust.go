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

// CratesVersionHtmlMarkdown
//
//	see: https://shields.io/badges/crates-io-3
//
// size: badges.MarkdownImgSizes
func CratesVersionHtmlMarkdown(repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, CratesVersion(repo), badges.MarkdownImgAlign, size, "crates.io version", "crates.io version")
	return fmt.Sprintf(
		"[%s](https://crates.io/crates/%s)",
		htmlContent, repo,
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

// CratesLicenseHtmlMarkdown
//
//	see: https://shields.io/badges/crates-io-1
//
// size: badges.MarkdownImgSizes
func CratesLicenseHtmlMarkdown(repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, CratesLicense(repo), badges.MarkdownImgAlign, size, "crates.io license", "crates.io license")
	return fmt.Sprintf(
		"[%s](https://crates.io/crates/%s)",
		htmlContent, repo,
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

// DocsRsHtmlMarkdown
//
//	see: https://shields.io/badges/docs-rs-1
//
// size: badges.MarkdownImgSizes
func DocsRsHtmlMarkdown(repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DocsRs(repo), badges.MarkdownImgAlign, size, "doc.rs", "doc.rs")
	return fmt.Sprintf(
		"[%s](https://docs.rs/%s)",
		htmlContent, repo,
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

// CratesDownloadLatestHtmlMarkdown
//
//	see: https://shields.io/badges/crates-io-latest
//
// size: badges.MarkdownImgSizes
func CratesDownloadLatestHtmlMarkdown(repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, CratesDownloadLatest(repo), badges.MarkdownImgAlign, size, "crates.io download latest", "crates.io download latest")
	return fmt.Sprintf(
		"[%s](https://crates.io/crates/%s)",
		htmlContent, repo,
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

// DepsRsCrateLatestHtmlMarkdown
//
//	see: https://deps.rs/
//
// size: badges.MarkdownImgSizes
func DepsRsCrateLatestHtmlMarkdown(repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DepsRsCrateLatest(repo), badges.MarkdownImgAlign, size, "deps.rs dependency", "deps.rs dependency")
	return fmt.Sprintf(
		"[%s](https://deps.rs/repo/github/%s/latest)",
		htmlContent, repo)
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

// DepsRsGithubHtmlMarkdown
//
//	see https://deps.rs/
//
// size: badges.MarkdownImgSizes
func DepsRsGithubHtmlMarkdown(user, repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DepsRsGithub(user, repo), badges.MarkdownImgAlign, size, "deps.rs dependency", "deps.rs dependency")
	return fmt.Sprintf(
		"[%s](https://deps.rs/repo/github/%s/%s)",
		htmlContent, user, repo)
}
