// Excercism excercise atbash-cipher
package atbash

import (
	"strings"
	"unicode"
)

// Encode a plain-text string into a cipher-text string using the atbash cipher
func Atbash(plain string) string {
	plain = normalize(plain)
	var cipherBuilder strings.Builder
	for i, r := range plain {
		if i != 0 && i%5 == 0 {
			cipherBuilder.WriteRune(' ')
		}
		cipherBuilder.WriteRune(encrypt(r))
	}
	return cipherBuilder.String()
}

// Normalize a string remove all but letters and digits and converting to lower case.
func normalize(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return strings.ToLower(builder.String())
}

// Encrypt a single rune using the atbash cipher.
func encrypt(r rune) rune {
	if unicode.IsDigit(r) {
		return r
	}
	return 'a' + ('z' - r)
}
