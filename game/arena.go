package game

type Arena struct {
	TL Coord
	BR Coord
}

func (a Arena) CollidingX(x int) bool {
	return x <= a.TL.X || x >= a.BR.X
}

func (a Arena) CollidingY(y int) bool {
	return y <= a.TL.Y || y >= a.BR.Y
}
