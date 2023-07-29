package utils

import "github.com/hajimehoshi/ebiten/v2"

const scale float64 = 1

type Sprite struct {
	X     int
	Y     int
	Image *ebiten.Image
}

func (sprite *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(sprite.X), float64(sprite.Y))
	options.GeoM.Scale(scale, scale)

	screen.DrawImage(sprite.Image, options)
}
