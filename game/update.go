package game

func (g *Game) Update() {
	g.playerLeft.Update()
	g.playerRight.Update()
	g.puck.Update()
}
