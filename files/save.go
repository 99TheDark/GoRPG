package files

import (
	"encoding/json"
	"game/top"
	"log"
)

func Save(world top.World) {
	var json_sprites []top.TileJSON
	for _, tile := range world {
		json_sprites = append(json_sprites, top.TileToJSON(*tile))
	}

	file, err := json.Marshal(json_sprites)
	if err != nil {
		log.Fatal(err)
	}

	WriteFile("res/save/map.json", file)
}
