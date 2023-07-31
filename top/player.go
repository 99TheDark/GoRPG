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
	if keys.Down("W") {
		p.Sprite.Y -= constants.PlayerSpeed
	}
	if keys.Down("S") {
		p.Sprite.Y += constants.PlayerSpeed
	}
	if keys.Down("A") {
		p.Sprite.X -= constants.PlayerSpeed
	}
	if keys.Down("D") {
		p.Sprite.X += constants.PlayerSpeed
	}

	utils.Clamp(&p.Sprite.X, 0, 7)
	utils.Clamp(&p.Sprite.Y, 0, 5)
}
