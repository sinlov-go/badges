package example

import (
	"github.com/sinlov-go/badges/npm_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNpm_badges(t *testing.T) {
	npmPackage := "git-tidier"

	assert.Equal(t, "https://img.shields.io/npm/v/git-tidier",
		npm_badges.VersionLatest(npmPackage),
	)
	assert.Equal(t, "[![npm version](https://img.shields.io/npm/v/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.VersionLatestMarkdown(npmPackage),
	)

	assert.Equal(t, "https://img.shields.io/node/v-lts/git-tidier",
		npm_badges.NodeLtsVersion(npmPackage),
	)

	assert.Equal(t, "[![node lts version](https://img.shields.io/node/v-lts/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.NodeLtsVersionMarkdown(npmPackage),
	)

	assert.Equal(t, "https://img.shields.io/npm/l/git-tidier",
		npm_badges.License(npmPackage),
	)

	assert.Equal(t, "[![npm license](https://img.shields.io/npm/l/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.LicenseMarkdown(npmPackage),
	)

	assert.Equal(t, "https://img.shields.io/npm/dt/git-tidier",
		npm_badges.DownloadLatestTimes(npmPackage),
	)

	assert.Equal(t, "[![npm download times](https://img.shields.io/npm/dt/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.DownloadLatestTimesMarkdown(npmPackage),
	)

	assert.Equal(t, "https://img.shields.io/npm/dm/git-tidier",
		npm_badges.DownloadLatestMonth(npmPackage),
	)

	assert.Equal(t, "[![npm download month](https://img.shields.io/npm/dm/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.DownloadLatestMonthMarkdown(npmPackage),
	)

	assert.Equal(t, "https://img.shields.io/npm/collaborators/git-tidier",
		npm_badges.Collaborators(npmPackage),
	)

	assert.Equal(t, "[![npm collaborators](https://img.shields.io/npm/collaborators/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.CollaboratorsMarkdown(npmPackage),
	)
}
