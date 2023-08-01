package top

import (
	"game/constants"
	"game/utils"
)

type World []*Tile

func (world *World) Add(x, y float64, path string) {
	new_tile := CreateTile(x, y, path)
	for _, tile := range *world {
		if tile == new_tile {
			return
		}
	}

	*world = append(*world, new_tile)
}

func (world *World) Solid(x, y float64) bool {
	for _, tile := range *world {
		if tile.Pos.X == x && tile.Pos.Y == y {
			return true
		}
	}
	return false
}

func (world *World) Blocked(pos *utils.Point, dir constants.Direction) bool {
	switch dir {
	case constants.Up:
		return world.Solid(pos.X, pos.Y-1)
	case constants.Down:
		return world.Solid(pos.X, pos.Y+1)
	case constants.Left:
		return world.Solid(pos.X-1, pos.Y)
	case constants.Right:
		return world.Solid(pos.X+1, pos.Y)
	case constants.UpLeft:
		return world.Solid(pos.X, pos.Y-1) && world.Solid(pos.X-1, pos.Y)
	case constants.DownLeft:
		return world.Solid(pos.X, pos.Y+1) && world.Solid(pos.X-1, pos.Y)
	case constants.UpRight:
		return world.Solid(pos.X, pos.Y-1) && world.Solid(pos.X+1, pos.Y)
	case constants.DownRight:
		return world.Solid(pos.X, pos.Y+1) && world.Solid(pos.X+1, pos.Y)
	}
	return true
}
