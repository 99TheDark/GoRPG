package main

import (
	"game/files"
	"game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var sprites utils.SpriteArray

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(sprites); i++ {
		sprites[i].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func init() {
	// sprites = files.Load()
	sprites.Add(0, 0, "grass.png")
	sprites.Add(1, 0, "grass.png")

	files.Save(sprites)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
