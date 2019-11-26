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

func (p *Position) RegionX() int {
	return (p.X >> 3) - 6;
}

func (p *Position) RegionY() int {
	return (p.Y >> 3) - 6;
}
