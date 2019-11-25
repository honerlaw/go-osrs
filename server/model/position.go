package model

type Position struct {
	X int
	Y int
	Z byte
}

func NewPosition(x int, y int, z byte) *Position {
	return &Position{
		X: x,
		Y: y,
		Z: z,
	}
}
