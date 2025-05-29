package game

import "github.com/RossLaing8417/ping-pong/ui"

type Game struct {
	Running      bool
	screen       *ui.Screen
}

func NewGame(screen *ui.Screen) *Game {
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
