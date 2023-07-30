package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	keys []ebiten.Key
}

func (k *Keyboard) Update() {
	k.keys = inpututil.PressedKeys()
}

func (k *Keyboard) Down(key string) bool {
	for i := 0; i < len(k.keys); i++ {
		if k.keys[i].String() == key {
			return true
		}
	}
	return false
}

func (k *Keyboard) Up(key string) bool {
	return !k.Down(key)
}
