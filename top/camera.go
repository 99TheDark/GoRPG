package top

type Camera struct {
	X      float64
	Y      float64
	Anchor Player
}

func CreateCamera(anchor Player) Camera {
	return Camera{anchor.Sprite.X, anchor.Sprite.Y, anchor}
}

func (cam *Camera) Update() {
	cam.X, cam.Y = cam.Anchor.Sprite.X, cam.Anchor.Sprite.Y
}
