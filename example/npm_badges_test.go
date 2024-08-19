package example

import (
	"github.com/sinlov-go/badges"
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

	assert.Equal(t, `[<img src="https://img.shields.io/npm/v/git-tidier" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="npm version" title="npm version">](https://www.npmjs.com/package/git-tidier)`,
		npm_badges.VersionLatestHtmlMarkdown(npmPackage, badges.MarkdownImgSizes))

	assert.Equal(t, "https://img.shields.io/node/v-lts/git-tidier",
		npm_badges.NodeLtsVersion(npmPackage),
	)

	assert.Equal(t, "[![node lts version](https://img.shields.io/node/v-lts/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.NodeLtsVersionMarkdown(npmPackage),
	)
	assert.Equal(t, `[<img src="https://img.shields.io/node/v-lts/git-tidier" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="node lts version" title="node lts version">](https://www.npmjs.com/package/git-tidier)`,
		npm_badges.NodeLtsVersionHtmlMarkdown(npmPackage, badges.MarkdownImgSizes))

	assert.Equal(t, "https://img.shields.io/npm/l/git-tidier",
		npm_badges.License(npmPackage),
	)

	assert.Equal(t, "[![npm license](https://img.shields.io/npm/l/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.LicenseMarkdown(npmPackage),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/npm/l/git-tidier" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="npm license" title="npm license">](https://www.npmjs.com/package/git-tidier)`,
		npm_badges.LicenseHtmlMarkdown(npmPackage, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/npm/dt/git-tidier",
		npm_badges.DownloadLatestTimes(npmPackage),
	)

	assert.Equal(t, "[![npm download times](https://img.shields.io/npm/dt/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.DownloadLatestTimesMarkdown(npmPackage),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/npm/dt/git-tidier" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="npm download times" title="npm download times">](https://www.npmjs.com/package/git-tidier)`,
		npm_badges.DownloadLatestTimesHtmlMarkdown(npmPackage, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/npm/dm/git-tidier",
		npm_badges.DownloadLatestMonth(npmPackage),
	)

	assert.Equal(t, "[![npm download month](https://img.shields.io/npm/dm/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.DownloadLatestMonthMarkdown(npmPackage),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/npm/dm/git-tidier" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="npm download month" title="npm download month">](https://www.npmjs.com/package/git-tidier)`,
		npm_badges.DownloadLatestMonthHtmlMarkdown(npmPackage, badges.MarkdownImgSizes),
	)

	assert.Equal(t, "https://img.shields.io/npm/collaborators/git-tidier",
		npm_badges.Collaborators(npmPackage),
	)

	assert.Equal(t, "[![npm collaborators](https://img.shields.io/npm/collaborators/git-tidier)](https://www.npmjs.com/package/git-tidier)",
		npm_badges.CollaboratorsMarkdown(npmPackage),
	)

	assert.Equal(t, `[<img src="https://img.shields.io/npm/collaborators/git-tidier" align=center sizes="(max-width: 500px) 100vw, (max-width: 900px) 50vw" alt="npm collaborators" title="npm collaborators">](https://www.npmjs.com/package/git-tidier)`,
		npm_badges.CollaboratorsHtmlMarkdown(npmPackage, badges.MarkdownImgSizes),
	)
}
