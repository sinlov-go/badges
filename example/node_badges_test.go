package example

import (
	"github.com/sinlov-go/badges/node_badges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Badges(t *testing.T) {
	gitUser := "bridgewwater"
	gitRepo := "git-tidier"
	assert.Equal(t,
		"https://img.shields.io/github/package-json/v/bridgewwater/git-tidier?label=package.json",
		node_badges.GitHubPackageJsonVersion(gitUser, gitRepo))

	assert.Equal(t, "[![github node package.json version](https://img.shields.io/github/package-json/v/bridgewwater/git-tidier?label=package.json)](https://github.com/bridgewwater/git-tidier)",
		node_badges.GitHubPackageJsonVersionMarkdown(gitUser, gitRepo))
}
