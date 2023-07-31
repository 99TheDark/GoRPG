package utils

import (
	"game/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	X      float64
	Y      float64
	Width  int
	Height int
	Path   string
	Image  *ebiten.Image
}

type SpriteJSON struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Path string  `json:"path"`
}

type SpriteArray []Sprite

func (sprite *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(sprite.X*constants.TileSizeFloat, (sprite.Y+1)*constants.TileSizeFloat-float64(sprite.Height))
	options.GeoM.Scale(constants.Scale, constants.Scale)

	screen.DrawImage(sprite.Image, options)
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

func SpriteToJSON(sprite Sprite) SpriteJSON {
	return SpriteJSON{sprite.X, sprite.Y, sprite.Path}
}

func SpriteFromJSON(json SpriteJSON) Sprite {
	return CreateSprite(json.X, json.Y, json.Path)
}

func CreateSprite(x, y float64, path string) Sprite {
	img := CreateImage(path)
	width, height := img.Size()

	return Sprite{x, y, width, height, path, img}
}
