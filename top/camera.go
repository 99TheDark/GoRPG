package top

type Camera struct {
	X      float64
	Y      float64
	Anchor Player
}

func CreateCamera(anchor Player) Camera {
	return Camera{anchor.Pos.X, anchor.Pos.Y, anchor}
}

func (cam *Camera) Update() {
	cam.X, cam.Y = cam.Anchor.Pos.X, cam.Anchor.Pos.Y
}
