package game

import "time"

// Game represents a current game state
type Game struct {
	Date time.Time `json:"date"`
}

// /----- UTILITY FUNCTIONS -----/

// getTodayDate returns today's date in 24h UTC format.
func getTodayDate() time.Time {
	return time.Now().UTC().Truncate(24 * time.Hour)
}
