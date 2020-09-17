// Exercism exercise crypto square.
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode a plain string into a ciphertext string using a square code.
func Encode(plain string) string {
	// The spec says to remove all punctuation. We interpret anything that is not a letter or digit
	// as punctuation.
	plain = normalize(plain)
	plainRunes := []rune(plain)
	plainLength := len(plainRunes)
	nrRows, nrColumns := rectangleSize(plainLength)
	var cipherBuilder strings.Builder
	for column := 0; column < nrColumns; column++ {
		for row := 0; row < nrRows; row++ {
			plainIndex := row*nrColumns + column
			var char rune
			if plainIndex >= plainLength {
				char = ' '
			} else {
				char = plainRunes[plainIndex]
			}
			cipherBuilder.WriteRune(char)
		}
		if column != nrColumns-1 {
			cipherBuilder.WriteRune(' ')
		}
	}
	return cipherBuilder.String()
}

func normalize(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return strings.ToLower(builder.String())
}

func rectangleSize(length int) (rows int, columns int) {
	r := int(math.Ceil(math.Sqrt(float64(length))))
	if (r-1)*r >= length {
		return r - 1, r
	}
	return r, r
}
