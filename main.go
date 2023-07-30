package main

import (
	"game/files"
	"game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var sprites utils.SpriteArray

type Game struct {
	Keys []ebiten.Key
}

func (g *Game) Update() error {
	g.Keys = inpututil.PressedKeys()
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
	w, h := 8, 6
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img := "grass.png"
			if x == 0 {
				if y == 0 {
					img = "grass_top_left.png"
				} else if y == h-1 {
					img = "grass_bottom_left.png"
				} else {
					img = "grass_left.png"
				}
			} else if x == w-1 {
				if y == 0 {
					img = "grass_top_right.png"
				} else if y == h-1 {
					img = "grass_bottom_right.png"
				} else {
					img = "grass_right.png"
				}
			} else if y == 0 {
				img = "grass_top.png"
			} else if y == h-1 {
				img = "grass_bottom.png"
			}

			sprites.Add(x, y, img)
		}
	}

	files.Save(sprites)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
