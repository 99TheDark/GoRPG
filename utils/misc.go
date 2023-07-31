package utils

import (
	"game/constants"
	"math"
)

func Clamp(num *float64, min, max float64) {
	*num = math.Min(math.Max(*num, min), max)
}

func Half(num int) float64 {
	return float64(num) / 2
}

func HalfScaled(num int) float64 {
	return float64(num) / (constants.Scale * 2)
}

func IndexOf[T comparable](arr []T, el T) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == el {
			return i
		}
	}
	return -1
}

func Remove[T comparable](arr []T, idx int) []T {
	return append(arr[:idx], arr[idx+1:]...)
}
