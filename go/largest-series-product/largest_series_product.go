// Excercism largest-series-product exercise.
package lsproduct

import (
	"errors"
	"fmt"
	"unicode"
)

// Compute largest product of all sub-strings of `span` digits in the given `digits` string.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return 0, errors.New("Span must be non-negative")
	}
	if span == 0 {
		// Not well defined, but expected by test cases
		return 1, nil
	}
	if span > len(digits) {
		return 0, errors.New("Span must be smaller than digits length")
	}
	maxProduct := 0
	for _, digit := range digits {
		if !unicode.IsDigit(digit) {
			return 0, errors.New(fmt.Sprintf("Encountered non-digit %v", digit))
		}
	}
	for i := 0; i <= len(digits)-span; i++ {
		product := 1
		for _, digit := range digits[i : i+span] {
			product *= int(digit - '0')
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}
