package top

import (
	"game/constants"
	"game/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Sprite    *utils.Sprite
	Direction constants.Direction
	Moving    bool
}

func CreatePlayer(x, y float64) Player {
	return Player{utils.CreateSprite(x, y, "character.png"), constants.Down, false}
}

func (p *Player) Update(keys *Keyboard) {
	if !p.Moving {
		dir, move := constants.NoDirection, false

		for _, key := range keys.Keys {
			switch key {
			case ebiten.KeyW:
				dir, move = dir.Combine(constants.Up), true
			case ebiten.KeyS:
				dir, move = dir.Combine(constants.Down), true
			case ebiten.KeyA:
				dir, move = dir.Combine(constants.Left), true
			case ebiten.KeyD:
				dir, move = dir.Combine(constants.Right), true
			}
		}

		if move {
			p.Direction, p.Moving = dir, true
		}
	} else {

	}

	utils.Clamp(&p.Sprite.X, 0, 7)
	utils.Clamp(&p.Sprite.Y, 0, 5)
}
