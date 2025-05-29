package main

import (
	"log"
	"os"
	"time"

	"github.com/RossLaing8417/ping-pong/game"
	"github.com/gdamore/tcell/v2"
)

func main() {
	logFile, err := os.OpenFile("/tmp/ping-pong.log", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)

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
	style = style.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

	width, height := screen.Size()
	g := game.NewGame(width, height)

	log.Println("Running game...")

	for g.IsRunning() {
		log.Println(">>> START >>>")

		screen.Clear()

		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		default:
			g.HandleEvent(event)
		}

		g.Update()

		for _, command := range g.GetDrawCommands(style) {
			screen.SetContent(command.X, command.Y, command.Data, nil, command.Style)
		}

		screen.Show()

		log.Println("<<<  END  <<<")

		time.Sleep(40 * time.Millisecond)
	}

	log.Println("Shutting down...")
}
