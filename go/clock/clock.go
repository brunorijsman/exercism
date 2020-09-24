// Exercism exercise clock
package clock

import "fmt"

const minutesPerDay int = 24 * 60

// Type to represent a clock.
type Clock int // Minutes after midnight

// Clock constructor.
func New(hour, minute int) Clock {
	return Clock(normalize(hour*60 + minute))
}

// Convert a clock to a human-readable string.
func (clock Clock) String() string {
	hour := clock / 60
	minute := clock % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

// Add minutes to a clock.
func (clock Clock) Add(minutes int) Clock {
	return Clock(normalize(int(clock) + minutes))
}

// Subtrack minutes from a clock.
func (clock Clock) Subtract(minutes int) Clock {
	return Clock(normalize(int(clock) - minutes))
}

func normalize(minutes int) int {
	normalized := minutes % minutesPerDay
	if normalized < 0 {
		normalized += minutesPerDay
	}
	return normalized
}
