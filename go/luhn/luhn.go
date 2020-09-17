// Validate a string (e.g. a credit card number) using the Luhn algorithm.
package luhn

import (
	"strings"
	"unicode"
)

// Validate a string (e.g. a credit card number) using the Luhn algorithm.
func Valid(validated string) bool {
	validated = strings.ReplaceAll(validated, " ", "")
	if len(validated) <= 1 {
		return false
	}
	sum := 0
	// If even length, double digits at index 0, 2, 4, etc. i.e. index%2 == 0
	// If odd length, double digits at index 1, 3, 5, etc. i.e. index%2 == 1
	doubleMod := len(validated) % 2
	for pos, digitRune := range validated {
		if !unicode.IsDigit(digitRune) {
			return false
		}
		digit := int(digitRune - '0')
		if pos%2 == doubleMod {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
