package files

import (
	"encoding/json"
	"game/constants"
	"game/top"
	"log"
)

func Load() top.World {
	var json_sprites []top.TileJSON

	err := json.Unmarshal(ReadFile(constants.MapLocation), &json_sprites)
	if err != nil {
		log.Fatal(json_sprites)
	}

	sprites := top.World{}
	for _, json_sprite := range json_sprites {
		tile := top.TileFromJSON(json_sprite)
		sprites = append(sprites, &tile)
	}

	return sprites
}
