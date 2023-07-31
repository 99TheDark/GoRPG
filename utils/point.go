package utils

type Point struct {
	X float64
	Y float64
}

func CreatePoint(x, y float64) *Point {
	return &Point{x, y}
}
