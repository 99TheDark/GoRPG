package main

import (
	"game/files"
	"game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var sprites []*utils.Sprite

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
	// Why does it insist on this instead of utils.Sprite(0, 0, utils.CreateImage(…))?
	sprite := utils.Sprite{X: 0, Y: 0, ImagePath: utils.CreateImage("grass.png")}
	sprites = append(sprites, &sprite)

	files.Save(sprites)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
