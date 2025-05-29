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

	baseStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(baseStyle)
	altStyle := baseStyle.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

	width, height := screen.Size()
	g := game.NewGame(width, height)

	log.Println("Running game...")

	go func() {
		for g.IsRunning() {
			screen.Clear()

			g.Update()

			for _, command := range g.GetDrawCommands(baseStyle, altStyle) {
				screen.SetContent(command.X, command.Y, command.Data, nil, command.Style)
			}

			screen.Show()

			time.Sleep(50 * time.Millisecond)
		}
	}()

	for g.IsRunning() {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		default:
			g.HandleEvent(event)
		}
	}

	log.Println("Shutting down...")
	time.Sleep(100 * time.Millisecond)
}
