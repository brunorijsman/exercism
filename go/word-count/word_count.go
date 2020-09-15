// Count the frequency of words in a text.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency of words.
type Frequency map[string]int

// Count the frequency of words in a text.
func WordCount(text string) Frequency {
	wordCount := Frequency{}
	for _, word := range strings.FieldsFunc(text, isSeparator) {
		wordCount[normalize(word)]++
	}
	return wordCount
}

// Is a character a word separator?
func isSeparator(c rune) bool {
	return unicode.IsSpace(c) || c == ','
}

// Normalize a word: convert all letters to lowercase, eliminate all characters except for letters
// and single quotes, and then remove all leading and trailing single quotes.
func normalize(word string) string {
	var normalizedWord string
	for _, c := range strings.ToLower(word) {
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == '\'' {
			normalizedWord += string(c)
		}
	}
	normalizedWord = strings.Trim(normalizedWord, "'")
	return normalizedWord
}
