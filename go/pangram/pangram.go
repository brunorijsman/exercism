// Determine whether a string is a pangram i.e. all letters a-z used at least once.
package pangram

import (
	"unicode"
)

// Void type used to implement a set as a map.
type void struct{}

// The one and only void value.
var yes = void{}

// Determine whether a string is a pangram i.e. all letters a-z used at least once.
func IsPangram(word string) bool {
	seen := map[rune]void{}
	for _, char := range word {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') {
			continue
		}
		seen[unicode.ToLower(char)] = yes
	}
	seenAll := true
	for char := 'a'; seenAll && char <= 'z'; char++ {
		if _, present := seen[char]; !present {
			seenAll = false
		}
	}
	return seenAll
}
