// Excercism exercise grains
package grains

import "errors"

// Compute the number of grains on the given square.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("Invalid square")
	}
	var grains uint64 = 1
	for s := 1; s < square; s++ {
		grains *= 2
	}
	return grains, nil
}

// Compute the total number of grains on the chess board.
func Total() uint64 {
	// The alternative way is to simply compute 2**65-1 but since Go doesn't have a uint64 power
	// operator, just is almost just as fast but clearer.
	var total uint64 = 0
	var grains uint64 = 1
	for s := 1; s <= 64; s++ {
		total += grains
		grains *= 2
	}
	return total
}
