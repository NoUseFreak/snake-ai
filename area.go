package main

type Area struct {
	Height, Width int
	Berry Coordinate
}

func NewArea(height int, width int) Area {
	area := Area{}
	area.Height = height
	area.Width = width

	return area
}

func (a *Area) IsFree(c Coordinate) bool {
	return c.x > 0 && c.y > 0 && c.x < a.Width && c.y < a.Height
}

func (a *Area) PlaceBerry(snake Snake) {
	for {
		//newPosition := Coordinate{rand.Intn(a.Width-2)+1, rand.Intn(a.Height-2)+1}
		newPosition := Coordinate{int(a.Width/2), int(a.Height/2)+1}
		if snake.IsFree(newPosition) {
			a.Berry = newPosition
			return
		}
	}
}

