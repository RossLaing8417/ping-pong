package game

type Player struct {
	X      int
	Top    int
	Bottom int
	Delta  int
	Score  int
}

func NewPlayer(x, arenaHeight int) Player {
	midY := arenaHeight / 2
	return Player{
		X:      x,
		Top:    midY - 5,
		Bottom: midY + 5,
		Delta:  1,
		Score:  0,
	}
}

func (p Player) Colliding(c Coord) bool {
	return c.X == p.X && c.Y >= p.Top && c.Y <= p.Bottom
}

func (p *Player) MoveUp(a Arena) {
	if !a.CollidingY(p.Top - p.Delta) {
		p.Top -= p.Delta
		p.Bottom -= p.Delta
	}
}

func (p *Player) MoveDown(a Arena) {
	if !a.CollidingY(p.Bottom + p.Delta) {
		p.Top += p.Delta
		p.Bottom += p.Delta
	}
}

func (p *Player) Update() {
}
