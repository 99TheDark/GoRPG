package connections

import "game/utils"

type TileType string

const (
	Grass TileType = "grass"
	Crate TileType = "crate"
)

var SolidTiles = []TileType{
	Crate,
}

var ConnectableTiles = []TileType{
	Grass,
}

func (tiletype TileType) Solid() bool {
	return utils.IndexOf(SolidTiles, tiletype) != -1
}

func (tiletype TileType) Connectable() bool {
	return utils.IndexOf(ConnectableTiles, tiletype) != -1
}

func (tiletype TileType) String() string {
	return string(tiletype)
}
