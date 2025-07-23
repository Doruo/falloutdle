package strings

import "strings"

// NormalizeCharacterName converts in string the values "_" into space,
// for example: "Roger_Maxson" -> "Roger Maxson"
func NormalizeString(str string) string {
	return strings.ReplaceAll(str, "_", " ")
}

// UnNormalizeString converts in string space int "_", as opposite to NormalizeString,
// for example: "Roger Maxson" -> "Roger_Maxson"
func UnNormalizeString(str string) string {
	return strings.ReplaceAll(str, " ", "_")
}
