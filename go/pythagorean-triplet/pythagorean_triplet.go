// Compute Pythagorean triplets.
package pythagorean

import (
	"math"
)

type Triplet [3]int

// Return a list of Pythagorean triplets (a,b,c) where a<b<c.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for c := min + 2; c <= max; c++ { // +2 since a and b must be at least 1
		for b := min + 1; b < c; b++ { // +1 since a must be at least 1
			a := int(math.Sqrt(float64(c*c - b*b))) // Given a fixed b and c, a is fully determined
			if a > min && a < b && a*a+b*b == c*c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

// Return a list of Pythagorean triplets that add up to sum.
func Sum(sum int) []Triplet {
	sumTriplets := []Triplet{}
	for _, triplet := range Range(1, sum) {
		if triplet[0]+triplet[1]+triplet[2] == sum {
			sumTriplets = append(sumTriplets, triplet)
		}
	}
	// The test cases expect the triplets to be returned in a certain order (this is a bug in the
	// test case in my honest opinion - the assignment does not specify any required order)
	return reverse(sumTriplets)
}

// Reverse a slice of triplets.
func reverse(triplets []Triplet) []Triplet {
	reversedTriplets := []Triplet{}
	for i := len(triplets) - 1; i >= 0; i-- {
		reversedTriplets = append(reversedTriplets, triplets[i])
	}
	return reversedTriplets
}
