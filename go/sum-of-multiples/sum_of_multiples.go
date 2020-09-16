// Compute sum of multiples.
package summultiples

// Compute the sum of numbers in range [1..limit) that are divisible by at least one divisor.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d != 0 && i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
