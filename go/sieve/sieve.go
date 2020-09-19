// Exercism exercise sieve.
package sieve

import "math"

// Find the prime numbers up to and including `upto` using the sieve of Eratosthenes.
func Sieve(upto int) []int {
	var composite = make([]bool, upto+1)
	rootOfUpto := int(math.Sqrt(float64(upto + 1)))
	for i := 2; i <= rootOfUpto; i++ {
		if !composite[i] {
			for j := i * i; j <= upto; j += i {
				composite[j] = true
			}
		}
	}
	primes := []int{}
	for i := 2; i <= upto; i++ {
		if !composite[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
