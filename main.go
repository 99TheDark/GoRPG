package main

import (
	"game/constants"
	"game/files"
	"game/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Sprite utils.Sprite
}

type Game struct {
	Keys    []ebiten.Key
	Sprites utils.SpriteArray
	Player  Player
}

func (g *Game) Update() error {
	g.Keys = inpututil.PressedKeys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.Sprites); i++ {
		g.Sprites[i].Draw(screen)
	}
	g.Player.Sprite.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	setup(g)

	return constants.ScreenWidth, constants.ScreenHeight
}

func setup(g *Game) {
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

			g.Sprites.Add(x, y, img)
		}
	}

	files.Save(g.Sprites)
}

func main() {
	game := Game{}

	game.Player = Player{utils.CreateSprite(4, 3, "character.png")}

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
