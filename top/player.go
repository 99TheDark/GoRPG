package top

import (
	"game/constants"
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
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
		switch key {
		case ebiten.KeyW:
			dir = dir.Combine(constants.Up)
		case ebiten.KeyS:
			dir = dir.Combine(constants.Down)
		case ebiten.KeyA:
			dir = dir.Combine(constants.Left)
		case ebiten.KeyD:
			dir = dir.Combine(constants.Right)
		}
	}

	utils.Clamp(&p.Sprite.X, 0, 7)
	utils.Clamp(&p.Sprite.Y, 0, 5)
}
