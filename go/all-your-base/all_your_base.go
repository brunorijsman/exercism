// Excerism exercise all-your-base
package allyourbase

import "errors"

// Convert a number, given as a slide of digits, from an input base to an output base.
func ConvertToBase(inBase int, inDigits []int, outBase int) ([]int, error) {
	if inBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	// Assumes the value fits in an int, which is indeed the case for all test cases.
	value := 0
	for _, inDigit := range inDigits {
		if inDigit < 0 || inDigit >= inBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		value = value*inBase + inDigit
	}
	outDigits := []int{}
	for value > 0 {
		outDigits = append(outDigits, value%outBase)
		value /= outBase
	}
	if len(outDigits) == 0 {
		outDigits = []int{0}
	}
	return reverse(outDigits), nil
}

func reverse(digits []int) []int {
	last := len(digits) - 1
	for i := 0; i <= last/2; i++ {
		digits[i], digits[last-i] = digits[last-i], digits[i]
	}
	return digits
}
