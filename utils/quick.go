package utils

import "game/constants"

func Half(num int) float64 {
	return float64(num) / 2
}

func HalfScaled(num int) float64 {
	return float64(num) / (constants.Scale * 2)
}
