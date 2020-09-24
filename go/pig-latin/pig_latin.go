// Exercism exercise pig-latin.
package piglatin

import "strings"

// Translate an English sentence into Pig-Latin.
func Sentence(sentence string) string {
	sentence = strings.ToLower(sentence)
	pigLatinWords := []string{}
	for _, word := range strings.Fields(sentence) {
		pigLatinWords = append(pigLatinWords, wordToPigLatin(word))
	}
	return strings.Join(pigLatinWords, " ")
}

func wordToPigLatin(word string) string {
	wordRunes := []rune(word)
	// Rule 1: If a word begins with a vowel sound, add an "ay" sound to the end of the word.
	if isVowel(wordRunes[0], 0) ||
		string(wordRunes[0:2]) == "xr" ||
		string(wordRunes[0:2]) == "yt" {
		return word + "ay"
	}
	n := leadingConsonants(word)
	if n > 0 {
		consonants := string(word[0:n])
		rest := string(word[n:])
		// Rule 3: If a word starts with a consonant sound followed by "qu", move it to the end of
		// the word, and then add an "ay" sound to the end of the word.
		if consonants[len(consonants)-1] == 'q' && rest[0] == 'u' {
			consonants = string(consonants[0 : len(consonants)-1])
			rest = string(rest[1:])
			return rest + consonants + "quay"
		}
		// Rule 4: If a word contains a "y" after a consonant cluster or as the second letter in a
		// two letter word it makes a vowel sound (e.g. "rhythm" -> "ythmrhay", "my" -> "ymay").
		if n > 1 && consonants[len(consonants)-1] == 'y' {
			consonants := string(consonants[0 : len(consonants)-1])
			return rest + consonants + "ay"
		}
		// Rule 2: If a word begins with a consonant sound, move it to the end of the word and then
		// add an "ay" sound to the end of the word.
		return rest + consonants + "ay"
	}
	return word
}

func isVowel(r rune, pos int) bool {
	if r == 'y' && pos == 0 {
		return false
	}
	return strings.ContainsRune("aeiouy", r)
}

func leadingConsonants(word string) int {
	for i, r := range word {
		if isVowel(r, i) {
			return i
		}
	}
	return len(word)
}
