// Package anagram detects anagrams of a word.
package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

// Compute the frequency histogram of characters in word.
func Frequencies(word string) map[rune]int {
	frequencies := map[rune]int{}
	for _, char := range word {
		frequencies[unicode.ToLower(char)]++
	}
	return frequencies
}

// Given a word and a list of candidates, return the subset of candidates that are anagrams of the
// word.
func Detect(word string, candidates []string) []string {
	anagrams := []string{}
	wordFrequencies := Frequencies(word)
	for _, candidate := range candidates {
		if !strings.EqualFold(word, candidate) &&
			reflect.DeepEqual(Frequencies(candidate), wordFrequencies) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
