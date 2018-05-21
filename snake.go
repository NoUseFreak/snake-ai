package main

import (
	"errors"
)


type Snake struct {
	Points[] Coordinate
}

func (s *Snake) IsFree(c Coordinate) bool {
	for _, point := range s.Points {
		if point.x == c.x && point.y == c.y {
			return false
		}
	}
	return true
}

func (s *Snake) MoveUp(area *Area) error {
	return s.move(Coordinate{s.Points[0].x,s.Points[0].y-1}, area)
}

func (s *Snake) MoveDown(area *Area) error {
	return s.move(Coordinate{s.Points[0].x,s.Points[0].y+1}, area)
}

func (s *Snake) MoveLeft(area *Area) error {
	return s.move(Coordinate{s.Points[0].x-1,s.Points[0].y}, area)
}

func (s *Snake) MoveRight(area *Area) error {
	return s.move(Coordinate{s.Points[0].x+1,s.Points[0].y}, area)
}

func (s *Snake) move(c Coordinate, area *Area) error {
	if !s.IsFree(c) {
		return errors.New("ran into snake")
	}
	if !area.IsFree(c) {
		return errors.New("ran into wall")
	}
	if !area.Berry.Equals(c) {
		s.Points = s.Points[:len(s.Points)-1]
	} else {
		area.PlaceBerry(*s)
	}
	s.Points = append([]Coordinate{c}, s.Points...)

	return nil
}