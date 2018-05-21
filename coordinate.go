package main

type Coordinate struct {
	x, y int
}

func (c Coordinate) Equals(o Coordinate) bool {
	return c.x == o.x && c.y == o.y
}