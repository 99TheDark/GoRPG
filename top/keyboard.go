package top

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	Keys []ebiten.Key
}

func (keyboard *Keyboard) Update() {
	keyboard.Keys = inpututil.PressedKeys()
}

func (keyboard *Keyboard) Down(key string) bool {
	for i := 0; i < len(keyboard.Keys); i++ {
		if keyboard.Keys[i].String() == key {
			return true
		}
	}
	return false
}

func (keyboard *Keyboard) Up(key string) bool {
	return !keyboard.Down(key)
}
