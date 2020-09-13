// Package raindrops converts a number into a string that contains raindrop sounds.
package raindrops

import "fmt"

// Function Convert converts a number into a string that contains raindrop sounds.
func Convert(raindrops int) string {
	var result string
	if raindrops%3 == 0 {
		result += "Pling"
	}
	if raindrops%5 == 0 {
		result += "Plang"
	}
	if raindrops%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = fmt.Sprint(raindrops)
	}
	return result
}
