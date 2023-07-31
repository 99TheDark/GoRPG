package utils

import (
	"game/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Width  int
	Height int
	Path   string
	Image  *ebiten.Image
}

func (sprite *Sprite) Draw(screen *ebiten.Image, options ebiten.DrawImageOptions, x, y float64) {
	options.GeoM.Translate(x*constants.TileSizeFloat, (y+1)*constants.TileSizeFloat-float64(sprite.Height))
	options.GeoM.Scale(constants.Scale, constants.Scale)

	screen.DrawImage(sprite.Image, &options)
}

func CreateSprite(path string) *Sprite {
	img := CreateImage(path)
	width, height := img.Size()

	return &Sprite{width, height, path, img}
}
