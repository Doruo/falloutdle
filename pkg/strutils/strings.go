package strutils

import (
	"regexp"
	"strings"
)

// RemplaceSpace converts all string spaces into "_" and retuns the result.
// Rx: "Name Surname" -> "Name_Surname"
func RemplaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

// RemplaceUnderscore converts all string values "_" into space and retuns the result,
// as opposite to NormalizeString.
// Rx: "Name_Surname" -> "Name Surname"
func RemplaceUnderscore(str string) string {
	return strings.ReplaceAll(str, "_", " ")
}

// NormalizeString removes under parentheses and their content.
func NormalizeString(s string) string {

	re := regexp.MustCompile(`\s*\([^)]*\)\s*`)
	s = re.ReplaceAllString(s, "")

	s = strings.ToLower(strings.TrimSpace(s))

	re = regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	return s
}
