package top

import (
	"game/constants"
	"game/utils"
)

type Player struct {
	Sprite    *utils.Sprite
	Direction constants.Direction
	Walking   bool
}

func CreatePlayer(x, y float64) Player {
	return Player{utils.CreateSprite(x, y, "character.png"), constants.Down, false}
}

func (p *Player) Update(keys *Keyboard) {
	var dir constants.Direction = constants.NoDirection

	for _, key := range keys.Keys {
		switch key.String() {
		case "W":
			dir = dir.Combine(constants.Up)
		case "S":
			dir = dir.Combine(constants.Down)
		case "A":
			dir = dir.Combine(constants.Left)
		case "D":
			dir = dir.Combine(constants.Right)
		}
	}

	utils.Clamp(&p.Sprite.X, 0, 7)
	utils.Clamp(&p.Sprite.Y, 0, 5)
}
