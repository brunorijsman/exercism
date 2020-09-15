// Check whether a string is an isogram.
package isogram

import "unicode"

// Empty data type to implement a set.
type Void struct{}

// Check whether a string is an isogram.
func IsIsogram(input string) bool {
	seenChars := map[rune]Void{}
	for _, char := range input {
		if char == '-' || char == ' ' {
			continue
		}
		lowerChar := unicode.ToLower(char)
		if _, present := seenChars[lowerChar]; present {
			return false
		}
		seenChars[lowerChar] = Void{}
	}
	return true
}
