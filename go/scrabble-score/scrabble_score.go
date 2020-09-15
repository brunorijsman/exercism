// Package scrabble computes the scrabble score for a word.
package scrabble

import (
	"fmt"
	"unicode"
)

var letterScores = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

// Score computes the scrabble score for a word.
func Score(word string) int {
	wordScore := 0
	for _, c := range word {
		lowerC := unicode.ToLower(c)
		letterScore, present := letterScores[lowerC]
		if !present {
			// We cannot return an error, so panic.
			panic(fmt.Sprintf("Unknown character %c", c))
		}
		wordScore += letterScore
	}
	return wordScore
}
