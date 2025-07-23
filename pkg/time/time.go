package time

import "time"

// getTodayDate returns today's date in 24h UTC format.
func GetTodayDate() time.Time {
	return time.Now().UTC().Truncate(24 * time.Hour)
}
