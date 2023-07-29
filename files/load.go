package files

import (
	"encoding/json"
	"game/utils"
	"log"
)

func Load() []*utils.Sprite {
	var sprites []*utils.Sprite
	err := json.Unmarshal(ReadFile("res/save/map.json"), &sprites)
	if err != nil {
		log.Fatal(err)
	}
	return sprites
}
