package top

import (
	"game/connections"
	"game/constants"
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Pos       *utils.Point
	Sprite    *utils.Sprite
	Type      constants.TileType
	Variation connections.Variation
}

type TileJSON struct {
	X         float64               `json:"x"`
	Y         float64               `json:"y"`
	Type      constants.TileType    `json:"type"`
	Variation connections.Variation `json:"var"`
}

func CreateTile(x, y float64, tiletype constants.TileType, variation connections.Variation) *Tile {
	return &Tile{
		utils.CreatePoint(x, y),
		utils.CreateSprite(tiletype.String() + "/" + variation.String() + ".png"),
		tiletype,
		variation,
	}
}

func (t *Tile) Draw(screen *ebiten.Image, options ebiten.DrawImageOptions) {
	t.Sprite.Draw(screen, options, t.Pos.X, t.Pos.Y)
}

func TileToJSON(tile Tile) TileJSON {
	return TileJSON{tile.Pos.X, tile.Pos.Y, tile.Type, tile.Variation}
}

func TileFromJSON(json TileJSON) Tile {
	return *CreateTile(json.X, json.Y, json.Type, json.Variation)
}
