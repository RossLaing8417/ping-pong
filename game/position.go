package game

type Coord struct {
	X int
	Y int
}

func (c Coord) Move(d Coord) Coord {
	return Coord{c.X + d.X, c.Y + d.Y}
}

type Position struct {
	TL Coord
	BR Coord
}

func (p Position) Colliding(o Position) bool {
	return false
}
