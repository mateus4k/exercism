package darts

import "math"

func Score(x float64, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	points := 0

	switch {
	case distance <= 1:
		points = 10
	case distance <= 5:
		points = 5
	case distance <= 10:
		points = 1
	}

	return points
}
