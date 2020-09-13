// Package hamming computes the hamming distance (i.e. the number of differences) between two
// equal length strings.
package hamming

import "errors"

// Note 1: the test specification doesn't say whether the DNA strands are encoded as ASCII strings
// or as Unicode strings. This code assumes ASCII strings. If the code has to handle unicode strings
// we would have to convert the string to a rune slice first ([]Rune).

// Note 2: this code doesn't check whether the string is a valid DNA string. We could return an
// error if we encounter anything else than A, C, G, T characters.

// Function Distance computes the hamming distance (i.e. the number of differences) between two
// equal length strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strings must be equal length")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
