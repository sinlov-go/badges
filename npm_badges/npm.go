package npm_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// VersionLatest
// See: https://shields.io/badges/npm-3
func VersionLatest(pkg string) string {
	return fmt.Sprintf(
		"%s/npm/v/%s",
		badges.ShieldsUrl, pkg,
	)
}

// VersionLatestMarkdown
// See: https://shields.io/badges/npm-3
func VersionLatestMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm version](%s)](https://www.npmjs.com/package/%s)",
		VersionLatest(pkg), pkg,
	)
}

// VersionLatestHtmlMarkdown
//
//	see: https://shields.io/badges/npm-3
//
// size: badges.MarkdownImgSizes
func VersionLatestHtmlMarkdown(pkg, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, VersionLatest(pkg), badges.MarkdownImgAlign, size, "npm version", "npm version")
	return fmt.Sprintf(
		"[%s](https://www.npmjs.com/package/%s)",
		htmlContent, pkg,
	)
}

// NodeLtsVersion
// See: https://shields.io/badges/node-lts
func NodeLtsVersion(pkg string) string {
	return fmt.Sprintf(
		"%s/node/v-lts/%s",
		badges.ShieldsUrl, pkg,
	)
}

// NodeLtsVersionMarkdown
// See: https://shields.io/badges/node-lts
func NodeLtsVersionMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![node lts version](%s)](https://www.npmjs.com/package/%s)",
		NodeLtsVersion(pkg), pkg,
	)
}

// NodeLtsVersionHtmlMarkdown
//
//	see: https://shields.io/badges/node-lts
//
// size: badges.MarkdownImgSizes
func NodeLtsVersionHtmlMarkdown(pkg, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, NodeLtsVersion(pkg), badges.MarkdownImgAlign, size, "node lts version", "node lts version")
	return fmt.Sprintf(
		"[%s](https://www.npmjs.com/package/%s)",
		htmlContent, pkg,
	)
}

// License
// See: https://shields.io/badges/npm-2
func License(pkg string) string {
	return fmt.Sprintf(
		"%s/npm/l/%s",
		badges.ShieldsUrl, pkg,
	)
}

// LicenseMarkdown
// See: https://shields.io/badges/npm-2
func LicenseMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm license](%s)](https://www.npmjs.com/package/%s)",
		License(pkg), pkg,
	)
}

// LicenseHtmlMarkdown
//
//	see: https://shields.io/badges/npm-2
//
// size: badges.MarkdownImgSizes
func LicenseHtmlMarkdown(pkg, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, License(pkg), badges.MarkdownImgAlign, size, "npm license", "npm license")
	return fmt.Sprintf(
		"[%s](https://www.npmjs.com/package/%s)",
		htmlContent, pkg,
	)
}

// DownloadLatestTimes
// See: https://shields.io/badges/npm-1
func DownloadLatestTimes(pkg string) string {
	return fmt.Sprintf(
		"%s/npm/dt/%s",
		badges.ShieldsUrl, pkg,
	)
}

// DownloadLatestTimesMarkdown
// See: https://shields.io/badges/npm-1
func DownloadLatestTimesMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm download times](%s)](https://www.npmjs.com/package/%s)",
		DownloadLatestTimes(pkg), pkg,
	)
}

// DownloadLatestTimesHtmlMarkdown
//
//	see: https://shields.io/badges/npm-1
//
// size: badges.MarkdownImgSizes
func DownloadLatestTimesHtmlMarkdown(pkg, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DownloadLatestTimes(pkg), badges.MarkdownImgAlign, size, "npm download times", "npm download times")
	return fmt.Sprintf(
		"[%s](https://www.npmjs.com/package/%s)",
		htmlContent, pkg,
	)
}

// DownloadLatestMonth
// See: https://shields.io/badges/npm-1
func DownloadLatestMonth(pkg string) string {
	return fmt.Sprintf(
		"%s/npm/dm/%s",
		badges.ShieldsUrl, pkg,
	)
}

// DownloadLatestMonthMarkdown
// See: https://shields.io/badges/npm-1
func DownloadLatestMonthMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm download month](%s)](https://www.npmjs.com/package/%s)",
		DownloadLatestMonth(pkg), pkg,
	)
}

// DownloadLatestMonthHtmlMarkdown
//
//	see: https://shields.io/badges/npm-1
//
// size: badges.MarkdownImgSizes
func DownloadLatestMonthHtmlMarkdown(pkg, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, DownloadLatestMonth(pkg), badges.MarkdownImgAlign, size, "npm download month", "npm download month")
	return fmt.Sprintf(
		"[%s](https://www.npmjs.com/package/%s)",
		htmlContent, pkg,
	)
}

// Collaborators
// See: https://shields.io/badges/npm-collaborators
func Collaborators(pkg string) string {
	return fmt.Sprintf(
		"%s/npm/collaborators/%s",
		badges.ShieldsUrl, pkg,
	)
}

// CollaboratorsMarkdown
// See: https://shields.io/badges/npm-collaborators
func CollaboratorsMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm collaborators](%s)](https://www.npmjs.com/package/%s)",
		Collaborators(pkg), pkg,
	)
}

// CollaboratorsHtmlMarkdown
//
//	see: https://shields.io/badges/npm-collaborators
//
// size: badges.MarkdownImgSizes
func CollaboratorsHtmlMarkdown(pkg, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, Collaborators(pkg), badges.MarkdownImgAlign, size, "npm collaborators", "npm collaborators")
	return fmt.Sprintf(
		"[%s](https://www.npmjs.com/package/%s)",
		htmlContent, pkg,
	)
}
