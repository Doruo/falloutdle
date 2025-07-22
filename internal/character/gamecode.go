package character

import "strings"

// GameCode represents standardized Fallout franchise game codes
type GameCode string

const (
	FO1    GameCode = "FO1"
	FO2    GameCode = "FO2"
	FO3    GameCode = "FO3"
	FNV    GameCode = "FNV"
	FO4    GameCode = "FO4"
	FO76   GameCode = "FO76"
	FOS    GameCode = "FOS"    // Fallout Shelter
	FOSO   GameCode = "FOSO"   // Fallout Shelter Online
	FOSBR  GameCode = "FOSBR"  // Fallout Shelter
	FO76SD GameCode = "FO76SD" // Fallout 76 Steel Dawn
	FO76SR GameCode = "FO76SR" // Fallout 76 Steel Reign
	FOT    GameCode = "FOT"    // Fallout Tactics
	FOBOS  GameCode = "FOBOS"  // Fallout: Brotherhood of Steel
	FBGNC  GameCode = "FBGNC"  // Fallout Board Game: New California
	FOWW   GameCode = "FOWW"   // Fallout: Wasteland Warfare
)

// AllGameCodes is the full list of known Fallout game codes
var AllGameCodes = []GameCode{
	FO1, FO2, FO3, FNV, FO4, FO76,
	FOS, FOSO, FOSBR,
	FO76SD, FO76SR,
	FOT, FOBOS, FBGNC, FOWW,
}

// GameFullName returns the full name of a game from its code
func (g GameCode) GameFullName() string {
	gameNames := map[GameCode]string{
		FO1:    "Fallout",
		FO2:    "Fallout 2",
		FO3:    "Fallout 3",
		FNV:    "Fallout: New Vegas",
		FO4:    "Fallout 4",
		FO76:   "Fallout 76",
		FOS:    "Fallout Shelter",
		FOSO:   "Fallout Shelter Online",
		FOSBR:  "Fallout Shelter",
		FO76SD: "Fallout 76: Steel Dawn",
		FO76SR: "Fallout 76: Steel Reign",
		FOT:    "Fallout Tactics",
		FOBOS:  "Fallout: Brotherhood of Steel",
		FBGNC:  "Fallout Board Game: New California",
		FOWW:   "Fallout: Wasteland Warfare",
	}

	if name, exists := gameNames[g]; exists {
		return name
	}
	return string(g)
}

// NormalizeGameCodes converts comma-separated game codes to slice
func NormalizeGameCodes(gamesStr string) []string {
	if gamesStr == "" {
		return []string{}
	}

	games := strings.Split(gamesStr, ",")
	var normalized []string

	for _, game := range games {
		game = strings.TrimSpace(game)
		if game != "" {
			normalized = append(normalized, game)
		}
	}

	return normalized
}
