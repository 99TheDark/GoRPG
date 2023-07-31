package main

import (
	"game/constants"
	"game/files"
	"game/top"
	"game/utils"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Keys    top.Keyboard
	Sprites utils.SpriteArray
	Player  top.Player
	Camera  top.Camera
}

func (g *Game) Update() error {
	g.Keys.Update()
	g.Player.Update(&g.Keys)
	g.Camera.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-g.Camera.X*constants.TileSizeFloat, -g.Camera.Y*constants.TileSizeFloat)
	options.GeoM.Translate(
		utils.HalfScaled(constants.ScreenWidth)-utils.Half(g.Player.Sprite.Width),
		utils.HalfScaled(constants.ScreenHeight)-utils.Half(g.Player.Sprite.Height),
	)

	screen.Fill(color.RGBA{78, 150, 199, 255})

	for i := 0; i < len(g.Sprites); i++ {
		g.Sprites[i].Draw(screen, *options)
	}
	g.Player.Sprite.Draw(screen, *options)

	ebitenutil.DebugPrint(screen, strconv.FormatInt(int64(ebiten.ActualFPS()), 10))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func main() {
	sprites := files.Load()
	player := top.CreatePlayer(3, 3)
	camera := top.CreateCamera(player)

	game := Game{Sprites: sprites, Player: player, Camera: camera}

	// files.Save(game.Sprites)

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("GoRPG")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
