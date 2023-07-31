package utils

import (
	"game/constants"
	"math"
)

func Clamp(num *float64, min, max float64) bool {
	switch {
	case *num > max:
		*num = max
		return true
	case *num < min:
		*num = min
		return true
	default:
		return false
	}
}

func Half(num int) float64 {
	return float64(num) / 2
}

func HalfScaled(num int) float64 {
	return float64(num) / (constants.Scale * 2)
}

func Normalize(num float64) float64 {
	if num == 0 {
		return 0
	} else {
		return num / math.Abs(num)
	}
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
