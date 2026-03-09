package models

import (
	"regexp"
	"strings"
)

// suidRegexp removes special characters, keep alphanumeric only with underscore
var suidRegexp = regexp.MustCompile(`[^a-zA-Z0-9_\s]+`)

// generateSUID generated string application id from name
// 1. trims spaces from both sides
// 2. replaces all non-alphanumerics with underscores
// 3. converts to lower case
// 4. removes spaces
func generateSUID(name string) string {
	cleaned := suidRegexp.ReplaceAllString(strings.TrimSpace(name), "_")

	cleaned = strings.ToLower(cleaned)
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	return cleaned
}
