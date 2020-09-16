// Compute the nth prime.
package prime

import "math"

// Compute the nth prime.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return -1, false
	}
	// Determine the first n primes using the sieve of Eratosthenes
	// See https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
	upperBound := nthPrimeUpperBound(n)
	var sieve = make([]bool, upperBound+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false
	rootOfUpperBound := int(math.Sqrt(float64(upperBound)))
	for i := 2; i <= rootOfUpperBound; i++ {
		if sieve[i] {
			for j := i * i; j <= upperBound; j += i {
				sieve[j] = false
			}
		}
	}
	// Extract the nth prime.
	count := 0
	nthPrime := -1
	for i := 0; i <= upperBound; i++ {
		if sieve[i] {
			count++
			if count == n {
				nthPrime = i
				break
			}
		}
	}
	if nthPrime == -1 {
		panic("Internal error - upper bound was too low")
	}
	return nthPrime, true
}

func nthPrimeUpperBound(n int) int {
	// See https://en.wikipedia.org/wiki/Prime-counting_function#Inequalities Dusart(1999)
	var f float64
	if n < 6 {
		f = 6.0
	} else {
		f = float64(n)
	}
	return int(f * math.Log(f*math.Log(f)))
}
