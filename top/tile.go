package top

import (
	"game/connections"
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Pos       *utils.Point
	Sprite    *utils.Sprite
	Type      connections.TileType
	Variation connections.Variation
}

type TileJSON struct {
	X         float64                `json:"x"`
	Y         float64                `json:"y"`
	Type      connections.TileType   `json:"type"`
	Variation *connections.Variation `json:"var,omitempty"`
}

func CreateTile(x, y float64, tiletype connections.TileType, variation connections.Variation) *Tile {
	var path string
	if tiletype.Connectable() {
		path = tiletype.String() + "/" + variation.String()
	} else {
		path = "objects/" + tiletype.String()
	}

	return &Tile{
		utils.CreatePoint(x, y),
		utils.CreateSprite(path + ".png"),
		tiletype,
		variation,
	}
}

func (t *Tile) Draw(screen *ebiten.Image, options ebiten.DrawImageOptions) {
	t.Sprite.Draw(screen, options, t.Pos.X, t.Pos.Y)
}

func TileToJSON(tile Tile) TileJSON {
	var variation *connections.Variation
	if tile.Variation == connections.Default {
		variation = nil
	} else {
		variation = &tile.Variation
	}

	return TileJSON{tile.Pos.X, tile.Pos.Y, tile.Type, variation}
}

func TileFromJSON(json TileJSON) Tile {
	var variation connections.Variation
	if json.Variation == nil {
		variation = connections.Default
	} else {
		variation = *json.Variation
	}

	return *CreateTile(json.X, json.Y, json.Type, variation)
}
