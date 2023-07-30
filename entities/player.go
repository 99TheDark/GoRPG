package entities

import "game/utils"

type Player struct {
	Sprite utils.Sprite
}

func CreatePlayer(x int, y int) Player {
	return Player{utils.CreateSprite(x, y, "character.png")}
}
