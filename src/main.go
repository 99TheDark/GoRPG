package main

import (
	"fmt"
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
	// sprite := Sprite{0, 0, createImage("../res/grass.png")}

	// fmt.Println(sprite)
	fmt.Println(createImage("../res/grass.png"))
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
