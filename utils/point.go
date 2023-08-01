package utils

type Point struct {
	X float64
	Y float64
}

func CreatePoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Add(q Point) Point {
	return *CreatePoint(p.X+q.X, p.Y+q.Y)
}

func (p Point) Sub(q Point) Point {
	return *CreatePoint(p.X-q.X, p.Y-q.Y)
}
