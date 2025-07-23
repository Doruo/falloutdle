package strings

import "strings"

// NormalizeString converts in string space int "_",
// for example: "Roger Maxson" -> "Roger_Maxson"
func NormalizeString(str string) string {
	return strings.ReplaceAll(str, " ", "_")
}

// UnNormalizeString converts in string the values "_" into space,
// as opposite to NormalizeString,
// for example: "Roger_Maxson" -> "Roger Maxson"
func UnNormalizeString(str string) string {
	return strings.ReplaceAll(str, "_", " ")
}
