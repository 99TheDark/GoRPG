package main

import "github.com/hajimehoshi/ebiten/v2"

const scale float64 = 1

type Sprite struct {
	image *ebiten.Image
	x     int
	y     int
}

func (sprite *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(sprite.x), float64(sprite.y))
	options.GeoM.Scale(scale, scale)

	screen.DrawImage(sprite.image, options)
}
