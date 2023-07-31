package files

import (
	"encoding/json"
	"game/constants"
	"game/utils"
	"log"
)

func Load() utils.SpriteArray {
	var json_sprites []*utils.SpriteJSON

	err := json.Unmarshal(ReadFile(constants.MapLocation), &json_sprites)
	if err != nil {
		log.Fatal(json_sprites)
	}

	sprites := utils.SpriteArray{}
	for _, json_sprite := range json_sprites {
		s := utils.SpriteFromJSON(*json_sprite)
		sprites = append(sprites, s)
	}

	return sprites
}
