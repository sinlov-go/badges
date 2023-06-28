package npm_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// VersionLatestMarkdown
// see: https://shields.io/badges/npm-3
func VersionLatestMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm version](%s/npm/v/%s)](https://www.npmjs.com/package/%s)",
		badges.ShieldsUrl, pkg, pkg,
	)
}

// NodeLtsVersion
// see: https://shields.io/badges/node-lts
func NodeLtsVersion(pkg string) string {
	return fmt.Sprintf(
		"[![node lts version](%s/node/v-lts/%s)](https://www.npmjs.com/package/%s)",
		badges.ShieldsUrl, pkg, pkg,
	)
}

// LicenseMarkdown
// see: https://shields.io/badges/npm-2
func LicenseMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm license](%s/npm/l/%s)](https://www.npmjs.com/package/%s)",
		badges.ShieldsUrl, pkg, pkg,
	)
}

// DownloadLatestTimesMarkdown
// see: https://shields.io/badges/npm-1
func DownloadLatestTimesMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm download times](%s/npm/dt/%s)](https://www.npmjs.com/package/%s)",
		badges.ShieldsUrl, pkg, pkg,
	)
}

// DownloadLatestMonthMarkdown
// see: https://shields.io/badges/npm-1
func DownloadLatestMonthMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm download month](%s/npm/dm/%s)](https://www.npmjs.com/package/%s)",
		badges.ShieldsUrl, pkg, pkg,
	)
}

// CollaboratorsMarkdown
// see: https://shields.io/badges/npm-collaborators
func CollaboratorsMarkdown(pkg string) string {
	return fmt.Sprintf(
		"[![npm collaborators](%s/npm/collaborators/%s)](https://www.npmjs.com/package/%s)",
		badges.ShieldsUrl, pkg, pkg,
	)
}
