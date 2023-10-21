package main

import "math/rand"

type Position struct {
	Y, X int
}

type Snake struct {
	direction Direction
	body      []Position
}

type Food Position

type Direction int8

const (
	Right Direction = iota
	Down
	Up
	Left
)

func NewSnake(position Position) *Snake {
	return &Snake{
		body: []Position{position},
	}
}

type Game struct {
	screen *Screen
	snake  Snake
	food   Food
}

func NewGame(s *Screen) *Game {
	width, height := s.Size()
	initial := Position{
		height / 2,
		width / 2,
	}
	random := Position{
		rand.Intn(height) + 1,
		rand.Intn(width) + 1,
	}

	game := &Game{
		screen: s,
		snake:  *NewSnake(initial),
		food:   Food(random),
	}

	s.HideCursor()

	return game
}

func (g *Game) Display() {
	g.screen.Clear()

	for i, pos := range g.snake.body {
		g.screen.MoveTo(pos)

		if i == 0 {
			g.screen.Printf("O")
		} else {
			g.screen.Printf("o")
		}
	}

	g.screen.MoveTo(Position(g.food))
	g.screen.Printf("x")

	g.screen.Flush()
}
