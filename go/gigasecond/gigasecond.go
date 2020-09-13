// Package gigasecond adds one gigasecond to a given time.
package gigasecond

import "time"

// Function GigaSecond returns a time one gigaseconds after the given time t.
func AddGigasecond(t time.Time) time.Time {
	const oneGigasecond time.Duration = 1000000000 * time.Second
	return t.Add(oneGigasecond)
}
