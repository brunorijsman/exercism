// Exercism exercise perfect-numbers.
package perfect

import "errors"

// The classification of a number.
type Classification int

const (
	ClassificationAbundant = iota
	ClassificationDeficient
	ClassificationPerfect
)

var (
	ErrOnlyPositive = errors.New("Only positive numbers")
)

// Classify a number as abundant, deficient, or perfect based on its aliquot sum.
func Classify(number int64) (Classification, error) {
	// See https://www.geeksforgeeks.org/sum-of-all-proper-divisors-of-a-natural-number/
	if number <= 0 {
		return -1, ErrOnlyPositive
	}
	if number == 1 {
		return ClassificationDeficient, nil
	}
	var sum int64 = 1
	for i := int64(2); i*i < number; i++ {
		if number%i == 0 {
			sum += i + number/i
		}
	}
	switch {
	case sum > number:
		return ClassificationAbundant, nil
	case sum < number:
		return ClassificationDeficient, nil
	default:
		return ClassificationPerfect, nil
	}
}
