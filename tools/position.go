package tools


type Position struct {
	X, Y int
}

func NewPosition(x, y int) Position { return Position{x, y} }