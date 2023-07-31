package files

import (
	"encoding/json"
	"game/utils"
	"log"
)

func Load() utils.SpriteArray {
	var sprites utils.SpriteArray
	err := json.Unmarshal(ReadFile("res/save/map.json"), &sprites)
	if err != nil {
		log.Fatal(err)
	}

	return sprites
}
