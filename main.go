package main

import (
	"game/constants"
	"game/entities"
	"game/files"
	"game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Keys    utils.Keyboard
	Sprites utils.SpriteArray
	Player  entities.Player
}

func (g *Game) Update() error {
	g.Keys.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.Sprites); i++ {
		g.Sprites[i].Draw(screen)
	}
	g.Player.Sprite.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func main() {
	game := Game{Player: entities.CreatePlayer(4, 3)}

	game.Sprites = files.Load()
	// files.Save(game.Sprites)

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
