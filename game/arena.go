package game

import "github.com/gdamore/tcell/v2"

type Arena struct {
	Style tcell.Style
	TL    Coord
	BR    Coord
}

func (a Arena) CollidingX(x int) bool {
	return x <= a.TL.X || x >= a.BR.X
}

func (a Arena) CollidingY(y int) bool {
	return y <= a.TL.Y || y >= a.BR.Y
}
