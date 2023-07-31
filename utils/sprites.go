package utils

import (
	"game/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	X         float64
	Y         float64
	Width     int
	Height    int
	ImagePath string
}

type SpriteArray []Sprite

func (sprite *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(sprite.X*constants.TileSizeFloat, sprite.Y*constants.TileSizeFloat)
	options.GeoM.Scale(constants.Scale, constants.Scale)

	screen.DrawImage(Images[sprite.ImagePath], options)
}

func (sprites *SpriteArray) Add(x, y float64, path string) {
	new_sprite := CreateSprite(x, y, path)

	for i := 0; i < len(*sprites); i++ {
		if (*sprites)[i] == new_sprite {
			return
		}
	}
	*sprites = append(*sprites, new_sprite)
}

func CreateSprite(x, y float64, path string) Sprite {
	return Sprite{x, y, 0, 0, CreateImage(path)}
}
