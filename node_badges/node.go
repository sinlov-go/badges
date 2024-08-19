package node_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
)

// GitHubPackageJsonVersion
//
//	get version from package.json in github
//
// See: https://shields.io/badges/git-hub-package-json-version-subfolder-of-monorepo
func GitHubPackageJsonVersion(user, repo string) string {
	return fmt.Sprintf("%s/github/package-json/v/%s/%s?label=package.json",
		badges.ShieldsUrl, user, repo)
}

// GitHubPackageJsonVersionMarkdown
//
//	get version from package.json in github
//
// See: https://shields.io/badges/git-hub-package-json-version-subfolder-of-monorepo
func GitHubPackageJsonVersionMarkdown(user, repo string) string {
	return fmt.Sprintf("[![github node package.json version](%s)](https://github.com/%s/%s)",
		GitHubPackageJsonVersion(user, repo), user, repo)
}

// GitHubPackageJsonVersionHtmlMarkdown
//
//	get version from package.json in github
//
// size: badges.MarkdownImgSizes
func GitHubPackageJsonVersionHtmlMarkdown(user, repo, size string) string {
	htmlContent := fmt.Sprintf(badges.MarkdownImageFmt, GitHubPackageJsonVersion(user, repo), badges.MarkdownImgAlign, size, "github node package.json version", "github node package.json version")
	return fmt.Sprintf(
		"[%s](https://github.com/%s/%s)",
		htmlContent, user, repo)
}
