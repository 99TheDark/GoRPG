package constants

const (
	TileSize     int     = 16   // Size of each tile in pixels
	TileBuffer   int     = 2    // Number of tiles outside of the screen to be calculated / rendered
	Scale        float64 = 4    // Tile scale multiplier
	ScreenWidth  int     = 960  // Width of window
	ScreenHeight int     = 640  // Height of window
	PlayerSpeed  float64 = 0.05 // Speed of the player
)

// Casted
const (
	TileSizeFloat float64 = float64(TileSize)
)
