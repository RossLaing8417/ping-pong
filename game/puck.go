package game

type Puck struct {
	Position Coord
	Delta    Coord
}

func NewPuck(x, y int) Puck {
	return Puck{
		Position: Coord{x, y},
		Delta:    Coord{1, 1},
	}
}

func (p *Puck) Update(a Arena, pL, pR Player) {
	next := p.Position.Move(p.Delta)
	if pL.Colliding(next) || pR.Colliding(next) {
		p.Delta.X *= -1
	} else if a.CollidingX(next.X) {
		// TODO: Point
		p.Delta.X *= -1
	}
	if a.CollidingY(next.Y) {
		p.Delta.Y *= -1
	}
	p.Position = p.Position.Move(p.Delta)
}
