package time

import "time"

// Today returns today's date in 24h UTC format.
func Today() time.Time {
	return time.Now().UTC().Truncate(24 * time.Hour)
}
