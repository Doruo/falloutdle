package strings

import "strings"

// NormalizeString converts all string spaces into "_" and retuns the result.
// Rx: "Name Surname" -> "Name_Surname"
func NormalizeString(str string) string {
	return strings.ReplaceAll(str, " ", "_")
}

// UnnormalizeString converts all string values "_" into space and retuns the result,
// as opposite to NormalizeString.
// Rx: "Name_Surname" -> "Name Surname"
func UnnormalizeString(str string) string {
	return strings.ReplaceAll(str, "_", " ")
}
