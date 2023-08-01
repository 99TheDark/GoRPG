package constants

type TileType string

const (
	Grass TileType = "grass"
	Crate TileType = "crate"
)

func (tiletype TileType) Solid() bool {
	switch tiletype {
	case Crate:
		return true
	}
	return false
}

func (tiletype TileType) String() string {
	return string(tiletype)
}
