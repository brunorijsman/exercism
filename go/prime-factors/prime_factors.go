// Exercism exercise prime-factors.
package prime

// Compute prime factors of n using trial division method.
func Factors(n int64) []int64 {
	factors := []int64{}
	for candidate := int64(2); n > 1; candidate++ {
		for ; n%candidate == 0; n /= candidate {
			factors = append(factors, candidate)
		}
	}
	return factors
}
