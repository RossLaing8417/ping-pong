package game

import "github.com/gdamore/tcell/v2"

type Coord struct {
	X int
	Y int
}

type Position struct {
	TL Coord
	BR Coord
}

type Game struct {
	running      bool
	winningScore int
	arena        Position
	playerLeft   Player
	playerRight  Player
	puck         Puck
	buffer       []DrawCommand
}

func NewGame(width, height int) *Game {
	buffer := make([]DrawCommand, 0, (width * height))
	return &Game{
		running:      true,
		winningScore: 15,
		arena: Position{
			TL: Coord{0, 0},
			BR: Coord{width - 1, height - 1},
		},
		playerLeft:  NewPlayer(2, height),
		playerRight: NewPlayer(width-3, height),
		puck:        NewPuck(width-1, height-2),
		buffer:      buffer,
	}
}

func (g *Game) IsRunning() bool {
	return g.running
}

func (g *Game) HandleEvent(event tcell.Event) {
	switch e := event.(type) {
	case *tcell.EventKey:
		if e.Key() == tcell.KeyEsc {
			// TODO: Pause
		}
		if e.Key() == tcell.KeyCtrlC {
			g.running = false
		}
	}
}

func (g *Game) GameOver() bool {
	if g.playerLeft.Score > g.winningScore || g.playerRight.Score > g.winningScore {
		return true
	}
	return false
}
