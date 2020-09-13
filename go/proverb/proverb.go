// Package proverb prints a proverbs.
package proverb

import "fmt"

// Function proverb prints a proverb, where all but the last line contains a word from the provided
// list of words.
func Proverb(words []string) []string {
	nrWords := len(words)
	lines := make([]string, 0, nrWords)
	if nrWords == 0 {
		return lines
	}
	for index, word := range words[:nrWords-1] {
		nextWord := words[index+1]
		line := fmt.Sprintf("For want of a %s the %s was lost.", word, nextWord)
		lines = append(lines, line)
	}
	line := fmt.Sprintf("And all for the want of a %s.", words[0])
	lines = append(lines, line)
	return lines
}
