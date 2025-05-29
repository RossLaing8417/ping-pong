package main

import (
	"log"
	"time"

	"github.com/RossLaing8417/ping-pong/game"
	"github.com/RossLaing8417/ping-pong/ui"
)

func main() {
	screen, err := ui.NewScreen()
	if err != nil {
		log.Panicln(err)
	}
	defer screen.Destroy()
	g := game.NewGame(screen)

	for g.Running {
		screen.Clear()
		g.HandleEvent(screen.PollEvent())
		g.Update()
		screen.Render()
	}
}
