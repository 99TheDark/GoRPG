package constants

type Direction int

const (
	Up          Direction = -1
	Down        Direction = -2
	Left        Direction = -3
	Right       Direction = -4
	UpLeft      Direction = -5
	DownLeft    Direction = -6
	UpRight     Direction = -7
	DownRight   Direction = -8
	NoDirection Direction = -9
)

func (dir Direction) Combine(next Direction) Direction {
	if dir == NoDirection {
		return next
	} else if !dir.TwoDirectional() && dir.Opposite() != next {
		switch dir {
		case Up:
			if next == Left {
				return UpLeft
			} else {
				return UpRight
			}
		case Down:
			if next == Left {
				return DownLeft
			} else {
				return DownRight
			}
		case Left:
			if next == Up {
				return UpLeft
			} else {
				return DownLeft
			}
		case Right:
			if next == Up {
				return UpRight
			} else {
				return DownRight
			}
		}
	}

	return dir
}

func (dir Direction) TwoDirectional() bool {
	return dir == UpRight || dir == UpLeft || dir == DownRight || dir == DownLeft
}

func (dir Direction) Opposite() Direction {
	switch dir {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	case UpLeft:
		return DownRight
	case DownLeft:
		return UpRight
	case UpRight:
		return DownLeft
	case DownRight:
		return UpLeft
	}

	return NoDirection
}

func (dir Direction) Deconstruct() []Direction {
	if dir.TwoDirectional() {
		switch dir {
		case UpLeft:
			return []Direction{Up, Left}
		case DownLeft:
			return []Direction{Down, Left}
		case UpRight:
			return []Direction{Up, Right}
		case DownRight:
			return []Direction{Down, Right}
		}
	}
	return []Direction{dir}
}

func (dir Direction) String() string {
	switch dir {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	case UpLeft:
		return "up left"
	case DownLeft:
		return "down left"
	case UpRight:
		return "up right"
	case DownRight:
		return "down right"
	}

	return "none"
}
