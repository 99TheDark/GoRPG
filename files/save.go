package files

import (
	"encoding/json"
	"game/utils"
	"log"
)

func Save(sprites utils.SpriteArray) {
	var json_sprites []utils.SpriteJSON
	for _, s := range sprites {
		json_sprites = append(json_sprites, utils.SpriteToJSON(s))
	}

	file, err := json.Marshal(json_sprites)
	if err != nil {
		log.Fatal(err)
	}

	WriteFile("res/save/map.json", file)
}
