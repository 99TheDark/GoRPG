package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// images
var flower_img *ebiten.Image

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(50, 50)
	options.GeoM.Scale(0.5, 1)

	screen.DrawImage(flower_img, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func init() {
	var err error

	flower_img, err = ebitenutil.NewImageFromURL("https://png.pngtree.com/png-clipart/20230408/original/pngtree-transparent-watercolor-flowers-png-image_9036505.png") // jpg's not working for some reason?

	if err != nil { // crash if an error occurs
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
