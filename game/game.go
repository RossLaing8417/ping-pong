package game

import "github.com/gdamore/tcell/v2"

type Game struct {
	Running      bool
	screen       *ui.Screen
	screen       tcell.Screen
}

func NewGame(screen tcell.Screen) *Game {
	return &Game{
		Running:      true,
		screen: screen,
	}
}

func (g *Game) HandleEvent(event ui.Event) {
	switch e := event.(type) {
	case *ui.EventResize:
		g.screen.Sync()
	case *ui.EventKey:
		if e.Key() == ui.KeyCtrlC {
			g.Running = false
		}
	}
}
