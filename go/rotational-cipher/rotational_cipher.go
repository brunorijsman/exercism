// Excercism exercise rotiational-cipher
package rotationalcipher

import (
	"strings"
	"unicode"
)

// Encrypt a plain-text string in to a cipher-text using the rotational cipher with the given shift
// key.
func RotationalCipher(plainText string, shiftKey int) string {
	var builder strings.Builder
	for _, r := range plainText {
		if unicode.IsLetter(r) {
			var base rune
			if unicode.IsUpper(r) {
				base = 'A'
			} else {
				base = 'a'
			}
			r = base + rune((int(r-base)+shiftKey)%26)
		}
		builder.WriteRune(r)
	}
	return builder.String()
}
