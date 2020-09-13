package romannumerals

import (
	"errors"
	"fmt"
	"strings"
)

func digitToRomanNumeral(digit int, oneSymbol, fiveSymbol, tenSymbol string) string {
	switch {
	case digit <= 3:
		return strings.Repeat(oneSymbol, digit)
	case digit == 4:
		return oneSymbol + fiveSymbol
	case digit <= 8:
		return fiveSymbol + strings.Repeat(oneSymbol, digit-5)
	case digit == 9:
		return oneSymbol + tenSymbol
	}
	panic("Unexpected digit")
}

func ToRomanNumeral(nr int) (string, error) {
	if nr <= 0 || nr > 3000 {
		return "", errors.New(fmt.Sprintf("Number %d out of allowed range 1-3000", nr))
	}
	roman := digitToRomanNumeral(nr/1000, "M", "", "")
	nr %= 1000
	roman += digitToRomanNumeral(nr/100, "C", "D", "M")
	nr %= 100
	roman += digitToRomanNumeral(nr/10, "X", "L", "C")
	nr %= 10
	roman += digitToRomanNumeral(nr, "I", "V", "X")
	return roman, nil
}
