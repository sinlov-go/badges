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
