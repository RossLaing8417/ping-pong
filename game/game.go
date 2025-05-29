package game

import "github.com/gdamore/tcell/v2"

type Game struct {
	running      bool
	winningScore int
	arena        Arena
	playerLeft   Player
	playerRight  Player
	puck         Puck
	buffer       []DrawCommand
}

func NewGame(width, height int) *Game {
	a := Arena{
		TL: Coord{0, 0},
		BR: Coord{width - 1, height - 1},
	}
	return &Game{
		running:      true,
		winningScore: 3,
		arena:        a,
		playerLeft:   NewPlayer(a.TL.X+2, height-2),
		playerRight:  NewPlayer(a.BR.X-2, height-2),
		puck:         NewPuck(a.BR.X/2, a.BR.Y/2),
		buffer:       make([]DrawCommand, 0, (width * height)),
	}
}

func (g *Game) IsRunning() bool {
	return g.running
}

func (g *Game) HandleEvent(event tcell.Event) {
	switch e := event.(type) {
	case *tcell.EventKey:
		if e.Key() == tcell.KeyCtrlC || e.Rune() == 'q' {
			g.running = false
		} else if e.Rune() == 'e' {
			g.playerLeft.MoveUp(g.arena)
		} else if e.Rune() == 'd' {
			g.playerLeft.MoveDown(g.arena)
		} else if e.Rune() == 'i' {
			g.playerRight.MoveUp(g.arena)
		} else if e.Rune() == 'k' {
			g.playerRight.MoveDown(g.arena)
		}
	}
}

func (g *Game) Update() {
	g.puck.Update(g.arena, &g.playerLeft, &g.playerRight)
	if g.gameOver() {
		g.running = false
	}
}

func (g *Game) gameOver() bool {
	if g.playerLeft.Score >= g.winningScore || g.playerRight.Score >= g.winningScore {
		return true
	}
	return false
}
