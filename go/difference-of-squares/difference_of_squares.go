// Compute square of sums, sum of squares, and the difference between the two.
package diffsquares

// Compute the square of the sum of the numbers [1...n].
func SquareOfSum(n int) int {
	sum := ((n * (n + 1)) / 2)
	return sum * sum
}

// Compute the sum of the squares of the numbers [1...n].
func SumOfSquares(n int) int {
	return (n * (n + 1) * ((n * 2) + 1)) / 6
}

// Compute the difference between square-of-sums and sum-of-squares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
