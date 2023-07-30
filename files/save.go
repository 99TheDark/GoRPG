package files

import (
	"encoding/json"
	"game/utils"
)

func Save(game_map utils.SpriteArray) {
	file, _ := json.Marshal(game_map)

	WriteFile("res/save/map.json", file)
}
