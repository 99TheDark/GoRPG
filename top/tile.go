package top

import (
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Pos    *utils.Point
	Sprite *utils.Sprite
}

type TileJSON struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Path string  `json:"path"`
}

type World []*Tile

func CreateTile(x, y float64, path string) *Tile {
	return &Tile{
		utils.CreatePoint(x, y),
		utils.CreateSprite(path),
	}
}

func (t *Tile) Draw(screen *ebiten.Image, options ebiten.DrawImageOptions) {
	t.Sprite.Draw(screen, options, t.Pos.X, t.Pos.Y)
}

func (world *World) Add(x, y float64, path string) {
	new_tile := CreateTile(x, y, path)
	for _, tile := range *world {
		if tile == new_tile {
			return
		}
	}

	*world = append(*world, new_tile)
}

func TileToJSON(tile Tile) TileJSON {
	return TileJSON{tile.Pos.X, tile.Pos.Y, tile.Sprite.Path}
}

func TileFromJSON(json TileJSON) Tile {
	return *CreateTile(json.X, json.Y, json.Path)
}
