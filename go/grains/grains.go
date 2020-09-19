// Excercism exercise grains
package grains

import "errors"

// Compute the number of grains on the given square.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("Invalid square")
	}
	return 1 << uint(square-1), nil
}

// Compute the total number of grains on the chess board.
func Total() uint64 {
	return ^uint64(0)
}
