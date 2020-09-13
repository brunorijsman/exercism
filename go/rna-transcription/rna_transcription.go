// Package strand computes the complementary RNA for a given DNA string.
package strand

import "strings"

// Note: the exercise does not specify whether the DNA string and the RNA string are ASCII encoded
// or Unicode encoded. This code assumes ASCII encoding. For Unicode encoding we would have to
// convert the string to a slice of Runes first.

// Note: having looked at other people's solutions, using string.NewReplacer and Replace seems to
// be the better way to go.

// Function ToRNA computes the complementary RNA for a given DNA string.
func ToRNA(dna string) string {
	length := len(dna)
	var rna strings.Builder
	rna.Grow(length)
	for _, char := range dna {
		switch char {
		case 'G':
			rna.WriteByte('C')
		case 'C':
			rna.WriteByte('G')
		case 'T':
			rna.WriteByte('A')
		case 'A':
			rna.WriteByte('U')
		default:
			// The exercise does not specify how to handle invalid strings, and the template
			// function doesn't allow us to return an error, so we panic.
			panic("Invalid letter in DNA")
		}
	}
	return rna.String()
}
