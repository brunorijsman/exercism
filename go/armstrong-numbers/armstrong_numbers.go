// Go exercise armstrong.
package armstrong

// Is number n an Armstrong number?
func IsNumber(n int) bool {
	// Compute the number of digits.
	nCopy := n
	numDigits := 0
	for nCopy > 0 {
		numDigits++
		nCopy /= 10
	}
	if numDigits == 0 {
		numDigits = 1
	}
	// Compute the sum of digits each raised to the power of the number of digits.
	nCopy = n
	sum := 0
	for nCopy > 0 {
		digit := nCopy % 10
		nCopy /= 10
		sum += power(digit, numDigits)
	}
	// Is it an Armstrong number?
	return sum == n
}

func power(n, p int) int {
	result := 1
	for i := 0; i < p; i++ {
		result *= n
	}
	return result
}
