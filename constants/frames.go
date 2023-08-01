package constants

// All non-completed animations will be the down animation for now.
var PlayerAnimation = map[Direction]([]string){
	Up: []string{
		"olivia/up_1.png",
		"olivia/up_2.png",
		"olivia/up_3.png",
	},
	Down: []string{
		"olivia/down_1.png",
		"olivia/down_2.png",
		"olivia/down_3.png",
	},
	Left: []string{
		"olivia/down_1.png",
		"olivia/down_2.png",
		"olivia/down_3.png",
	},
	Right: []string{
		"olivia/down_1.png",
		"olivia/down_2.png",
		"olivia/down_3.png",
	},
	UpLeft: []string{
		"olivia/up_left_1.png",
		"olivia/up_left_2.png",
		"olivia/up_left_3.png",
	},
	DownLeft: []string{
		"olivia/down_left_1.png",
		"olivia/down_left_2.png",
		"olivia/down_left_3.png",
	},
	UpRight: []string{
		"olivia/up_right_1.png",
		"olivia/up_right_2.png",
		"olivia/up_right_3.png",
	},
	DownRight: []string{
		"olivia/down_right_1.png",
		"olivia/down_right_2.png",
		"olivia/down_right_3.png",
	},
}
