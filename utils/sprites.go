package utils

import "github.com/hajimehoshi/ebiten/v2"

const scale float64 = 10

type Sprite struct {
	X         int
	Y         int
	ImagePath string
}

func (sprite *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(sprite.X), float64(sprite.Y))
	options.GeoM.Scale(scale, scale)

	screen.DrawImage(Images[sprite.ImagePath], options)
}
