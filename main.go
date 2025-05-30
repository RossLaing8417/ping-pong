package main

import (
	"log"
	"os"

	"github.com/RossLaing8417/ping-pong/game"
)

func main() {
	logFile, err := os.OpenFile("/tmp/ping-pong.log", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)

	g, err := game.NewGame()
	if err != nil {
		log.Fatalln(err)
	}
	defer g.End()

	log.Println("Running game...")

	g.Run()

	log.Println("Shutting down...")
}
