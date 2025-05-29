package game

type Player struct {
	Position Position
	Delta    int
	Score    int
}

func NewPlayer(x, height int) Player {
	midY := height / 2
	return Player{
		Position: Position{
			TL: Coord{x, midY - 5},
			BR: Coord{x, midY + 5},
		},
		Delta: 0,
		Score: 0,
	}
}

func (p *Player) Update() {
	p.Position.TL.Y += p.Delta
	p.Position.BR.Y += p.Delta
}
