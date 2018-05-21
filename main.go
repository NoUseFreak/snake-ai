package main

import (
	"github.com/labstack/gommon/log"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//RunGame()
	RunAI()

}

func RunGame() {

	game := NewGame(NewArea(20, 20))

	dirOptions := []string{
		"forward",
		"left",
		"right",
	}

	for {
		for i := 1; i <= 10; i++ {
			CallClear()
			game.Render()
			game.Debug()

			time.Sleep(100 * time.Millisecond)

			switch dirOptions[rand.Intn(len(dirOptions))] {
			case "forward":
				if err := game.GoForward(); err != nil {
					log.Fatal(err)
				}
			case "left":
				if err := game.GoLeft(); err != nil{
					log.Fatal(err)
				}
			case "right":
				if err := game.GoRight(); err != nil{
					log.Fatal(err)
				}
			}
		}
	}
}

