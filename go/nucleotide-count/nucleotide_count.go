// Package dna generates a histogram of nucleotide frequencies in a DNA strand.
package dna

import (
	"errors"
	"fmt"
	"strings"
)

type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (dna DNA) Counts() (Histogram, error) {
	var histogram Histogram = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range dna {
		if !strings.ContainsRune("ACGT", nucleotide) {
			return histogram, errors.New(fmt.Sprintf("Invald nucleotide %c", nucleotide))
		} else {
			histogram[nucleotide]++
		}
	}
	return histogram, nil
}
