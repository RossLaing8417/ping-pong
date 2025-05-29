package main

import (
	"log"
	"time"

	"github.com/RossLaing8417/ping-pong/game"
	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}
	if err = screen.Init(); err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Panicln(err)
	}
	defer screen.Fini()
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(style)
	g := game.NewGame(screen)

	for g.Running {
		screen.Clear()
		g.HandleEvent(screen.PollEvent())
		g.Update()

		screen.SetContent(0, 0, 'H', nil, style)
		screen.SetContent(1, 0, 'i', nil, style)
		screen.SetContent(2, 0, '!', nil, style)
		screen.Show()

		time.Sleep(40 * time.Millisecond)
	}
}
