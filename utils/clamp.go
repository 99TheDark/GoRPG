package utils

import "math"

func Clamp(num *float64, min, max float64) {
	*num = math.Min(math.Max(*num, min), max)
}
