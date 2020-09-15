package etl

import (
	"fmt"
	"strings"
)

type OldScoring map[int][]string
type NewScoring map[string]int

func Transform(oldScoring OldScoring) NewScoring {
	newScoring := make(NewScoring)
	for score, letters := range oldScoring {
		for _, letter := range letters {
			if len(letter) != 1 {
				// The test cases force us to represent letters as strings; it would have been
				// a better choice to represent letters as runes.
				panic(fmt.Sprintf("Multi-character letter %s in old scoring", letter))
			}
			lowerLetter := strings.ToLower(letter)
			_, alreadyPresent := newScoring[letter]
			if alreadyPresent {
				panic(fmt.Sprintf("Multiple scores for the same letter %s in old scoring", letter))
			}
			newScoring[lowerLetter] = score
		}
	}
	return newScoring
}
