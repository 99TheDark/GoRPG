package main

import (
	"fmt"
	"game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// var sprites []*Sprite

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func init() {
	// Why does it insist on this instead of utils.Sprite(0, 0, utils.CreateImage(…))?
	sprite := utils.Sprite{X: 0, Y: 0, Image: utils.CreateImage("res/grass.png")}

	fmt.Println(sprite)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
