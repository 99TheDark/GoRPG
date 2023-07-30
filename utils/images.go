package utils

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Images = map[string]*ebiten.Image{}

func CreateImage(path string) string {
	img, _, err := ebitenutil.NewImageFromFile("res/sprites/" + path)

	if err != nil {
		log.Fatal(err)
	}

	Images[path] = img
	return path
}
