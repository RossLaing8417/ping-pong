package game

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen       tcell.Screen
	running      bool
	winningScore int
	arena        Arena
	playerLeft   Player
	playerRight  Player
	puck         Puck
	buffer       []DrawCommand
	baseStyle    tcell.Style
}

func NewGame() (*Game, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err = screen.Init(); err != nil {
		return nil, err
	}

	baseStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(baseStyle)
	blockStyle := baseStyle.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

	width, height := screen.Size()
	a := Arena{
		Style: blockStyle,
		TL:    Coord{0, 0},
		BR:    Coord{width - 1, height - 1},
	}

	return &Game{
		screen:       screen,
		running:      true,
		winningScore: 3,
		arena:        a,
		playerLeft:   NewPlayer(blockStyle, a.TL.X+2, height-2),
		playerRight:  NewPlayer(blockStyle, a.BR.X-2, height-2),
		puck:         NewPuck(blockStyle, a.BR.X/2, a.BR.Y/2),
		buffer:       make([]DrawCommand, 0, (width * height)),
		baseStyle:    baseStyle,
	}, nil
}

func (g *Game) Run() {
	go g.loop()

	for g.running {
		switch event := g.screen.PollEvent().(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC || event.Rune() == 'q' {
				g.running = false
			} else if event.Rune() == 'e' {
				g.playerLeft.MoveUp(g.arena)
			} else if event.Rune() == 'd' {
				g.playerLeft.MoveDown(g.arena)
			} else if event.Rune() == 'i' {
				g.playerRight.MoveUp(g.arena)
			} else if event.Rune() == 'k' {
				g.playerRight.MoveDown(g.arena)
			}
		}
	}
}

func (g *Game) loop() {
	for g.running {
		g.screen.Clear()

		g.Update()

		for _, command := range g.GetDrawCommands() {
			g.screen.SetContent(command.X, command.Y, command.Data, nil, command.Style)
		}

		g.screen.Show()

		time.Sleep(50 * time.Millisecond)
	}
}

func (g *Game) End() {
	g.screen.Fini()
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
