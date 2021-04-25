package triangle

import "math"

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	// All sides must be positive integers, and not infinite.
	for _, s := range [...]float64{a, b, c} {
		if s <= 0 || math.IsNaN(s) || math.IsInf(s, 0) {
			return NaT
		}
	}

	// The triangle is impossible if the sum of the length of any two sides is less than the length of the third side.
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// If all sides are equal in length, the triangle is equilateral.
	if a == b && b == c {
		return Equ
	}

	// If two sides are equal in length, the triagle is isoceles.
	if a == b || a == c || b == c {
		return Iso
	}

	// If all three sides are different lengths the triangle is scalene.
	return Sca
}
