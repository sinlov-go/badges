package shields_badges

import (
	"fmt"
	"github.com/sinlov-go/badges"
	"net/url"
	"strings"
)

// StaticBadgeBlue
// See: https://shields.io/badges/static-badge
func StaticBadgeBlue(label, content string) string {
	return StaticBadge(label, content, "blue")
}

// StaticBadgeGreen
// See: https://shields.io/badges/static-badge
func StaticBadgeGreen(label, content string) string {
	return StaticBadge(label, content, "green")
}

// StaticBadgeYellow
// See: https://shields.io/badges/static-badge
func StaticBadgeYellow(label, content string) string {
	return StaticBadge(label, content, "yellow")
}

// StaticBadgeRed
// See: https://shields.io/badges/static-badge
func StaticBadgeRed(label, content string) string {
	return StaticBadge(label, content, "red")
}

// StaticBadgeOrange
// See: https://shields.io/badges/static-badge
func StaticBadgeOrange(label, content string) string {
	return StaticBadge(label, content, "orange")
}

// StaticBadgePurple
// See: https://shields.io/badges/static-badge
func StaticBadgePurple(label, content string) string {
	return StaticBadge(label, content, "purple")
}

// StaticBadge
// See: https://shields.io/badges/static-badge
func StaticBadge(label string, content string, color string) string {
	label = strings.ReplaceAll(label, "-", "--")
	label = strings.ReplaceAll(label, "_", "__")
	label = strings.ReplaceAll(label, " ", "_")
	content = strings.ReplaceAll(content, "-", "--")
	content = strings.ReplaceAll(content, "_", "__")
	content = strings.ReplaceAll(content, " ", "_")
	urlPath := fmt.Sprintf("%s-%s-%s", label, content, color)
	queryEscape := url.QueryEscape(urlPath)
	return fmt.Sprintf("%s/badge/%s",
		badges.ShieldsUrl, queryEscape)
}
