package game

type Puck struct {
	Position Coord
	Delta    Coord
}

func NewPuck(width, height int) Puck {
	return Puck{
		Position: Coord{width / 2, height / 2},
		Delta:    Coord{1, 1},
	}
}

func (p *Puck) Update() {
	p.Position.X += p.Delta.X
	p.Position.Y += p.Delta.Y
}
