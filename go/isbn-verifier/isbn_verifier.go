// Check validity of an ISBN number.
package isbn

import "unicode"

// Check validity of an ISBN number.
func IsValidISBN(isbn string) bool {
	isbnAsRunes := []rune(isbn)
	var dashes []int
	var digits []int
	var check int
	switch len(isbnAsRunes) {
	case 13:
		dashes = []int{1, 5, 11}
		digits = []int{0, 2, 3, 4, 6, 7, 8, 9, 10}
		check = 12
	case 10:
		dashes = []int{}
		digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		check = 9
	default:
		return false
	}
	for _, pos := range dashes {
		if isbnAsRunes[pos] != '-' {
			return false
		}
	}
	checksum := 0
	for i, pos := range digits {
		digit := isbnAsRunes[pos]
		if !unicode.IsDigit(digit) {
			return false
		}
		checksum += (i + 1) * int(digit-'0')
	}
	checksum %= 11
	var checksumRune rune
	if checksum == 10 {
		checksumRune = 'X'
	} else {
		checksumRune = rune(int('0') + checksum)
	}
	return isbnAsRunes[check] == checksumRune
}
