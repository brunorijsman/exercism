// Package triage categorizes triangles.
package triangle

import (
	"math"
)

type Kind int

const (
	NaT Kind = iota // Not a triangle
	Deg             // Degenerate
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// Function isBad returns true if x is a bad number (NaN, positive Inf, or negative Inf)
func isBad(x float64) bool {
	return math.IsNaN(x) || math.IsInf(x, 0)
}

// Function KindFromSides returns the kind of triangle given its side-lengths a, b, and c.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isBad(a) || isBad(b) || isBad(c):
		return NaT
	case a <= 0.0 || b <= 0.0 || c <= 0.0:
		return NaT
	case a < b+c || b < a+c || c < a+b:
		return NaT
	case a == b+c || b == a+c || c == a+b:
		return Deg
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
