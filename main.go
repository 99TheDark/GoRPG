package main

import (
	"game/constants"
	"game/entities"
	"game/files"
	"game/utils"
	"image/color"
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
	g.Player.Update(&g.Keys)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{78, 150, 199, 255})

	for i := 0; i < len(g.Sprites); i++ {
		g.Sprites[i].Draw(screen)
	}
	g.Player.Sprite.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func main() {
	game := Game{Sprites: files.Load(), Player: entities.CreatePlayer(4, 3)}

	// files.Save(game.Sprites)

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("GoRPG")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
