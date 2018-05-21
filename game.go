package main

import (
	"fmt"
	"errors"
	"math"
)

type Game struct {
	area     Area
	snake    Snake
	lastMove string
	startLen int
	stepCount int
}

const (
	GAME_UP = "up"
	GAME_DOWN = "down"
	GAME_LEFT = "left"
	GAME_RIGHT = "right"
)

func NewGame(area Area) Game {
	game := Game{}
	game.startLen = 3
	game.area = area
	game.stepCount = 0
	snakeCoords := []Coordinate{}
	for i := 0; i < game.startLen; i++ {
		snakeCoords = append(snakeCoords, Coordinate{int(area.Width/2),int(area.Height/2) - i})
	}
	game.snake = Snake{snakeCoords}
	game.lastMove = GAME_DOWN

	game.area.PlaceBerry(game.snake)

	return game
}

func (g *Game) Render() {
	fmt.Print("\n")
	for y := 0; y <= g.area.Height; y++ {
		for x := 0; x <= g.area.Width; x++ {
			fmt.Print(" ")
			newCoordinate := Coordinate{x, y}
			if !g.area.IsFree(newCoordinate) {
				fmt.Print("X")
			} else if !g.snake.IsFree(newCoordinate) {
				fmt.Print("*")
			} else if g.area.Berry.Equals(newCoordinate) {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (g *Game) MoveDirection(direction string) error {
	g.stepCount++
	switch direction {
	case GAME_UP:
		return g.MoveUp()
	case GAME_DOWN:
		return g.MoveDown()
	case GAME_LEFT:
		return g.MoveLeft()
	case GAME_RIGHT:
		return g.MoveRight()
	}
	return errors.New("invalid direction")
}

func (g *Game) MoveUp() error {
	if g.lastMove == GAME_DOWN {
		return g.snake.MoveDown(&g.area)
	}
	g.lastMove = GAME_UP
	return g.snake.MoveUp(&g.area)
}

func (g *Game) MoveDown() error {
	if g.lastMove == GAME_UP {
		return g.snake.MoveUp(&g.area)
	}
	g.lastMove = GAME_DOWN
	return g.snake.MoveDown(&g.area)
}

func (g *Game) MoveLeft() error {
	if g.lastMove == GAME_RIGHT {
		return g.snake.MoveRight(&g.area)
	}
	g.lastMove = GAME_LEFT
	return g.snake.MoveLeft(&g.area)
}

func (g *Game) MoveRight() error {
	if g.lastMove == GAME_LEFT {
		return g.snake.MoveLeft(&g.area)
	}
	g.lastMove = GAME_RIGHT
	return g.snake.MoveRight(&g.area)
}

func (g *Game) GoForward() error {
	return g.MoveDirection(g.lastMove)
}

func (g *Game) GoLeft() error {
	switch g.lastMove {
	case GAME_UP:
		return g.MoveDirection(GAME_LEFT)
	case GAME_DOWN:
		return g.MoveDirection(GAME_RIGHT)
	case GAME_LEFT:
		return g.MoveDirection(GAME_DOWN)
	case GAME_RIGHT:
		return g.MoveDirection(GAME_UP)
	}
	return errors.New(fmt.Sprintf("GoLeft failed, %s", g.lastMove))
}

func (g *Game) GoRight() error {
	switch g.lastMove {
	case GAME_UP:
		return g.MoveDirection(GAME_RIGHT)
	case GAME_DOWN:
		return g.MoveDirection(GAME_LEFT)
	case GAME_LEFT:
		return g.MoveDirection(GAME_UP)
	case GAME_RIGHT:
		return g.MoveDirection(GAME_DOWN)
	}
	return errors.New(fmt.Sprintf("GoRight failed, %s", g.lastMove))
}

func (g *Game) GetScore() float64 {
	maxDistance := math.Sqrt(math.Pow(float64(g.area.Height-2), 2) + math.Pow(float64(g.area.Width-2), 2))
	return maxDistance * float64(g.GetLevel() + 1) + g.GetDistanceScore() - float64(g.stepCount) * 0.1
}

func (g *Game) GetLevel() int {
	return len(g.snake.Points) - g.startLen + 1
}

func (g *Game) GetDistance() float64 {
	head := g.snake.Points[0]
	berry := g.area.Berry

	return math.Sqrt(math.Pow(float64(head.x - berry.x), 2) + math.Pow(float64(head.y - berry.y), 2))
}

func (g *Game) GetDistanceScore() float64 {
	return math.Sqrt(math.Pow(float64(g.area.Width - 2), 2) + math.Pow(float64(g.area.Height - 2),2)) - g.GetDistance()
}

func (g *Game) Reset() {

}

func (g *Game) Debug() {
	fmt.Printf("Level:    %d\n", g.GetLevel())
	fmt.Printf("Steps:    %d\n", g.stepCount)
	fmt.Printf("Score:    %f\n", g.GetScore())
	fmt.Printf("Distance: %f\n", g.GetDistanceScore())
	fmt.Println(g)
}