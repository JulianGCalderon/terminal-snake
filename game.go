package main

import (
	"io"
	"math/rand"
	"time"
)

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

func (s *Snake) Head() Position {
	return s.body[0]
}

func (s *Snake) Advance() {
	for i := len(s.body) - 1; i > 0; i-- {
		s.body[i] = s.body[i-1]
	}

	s.body[0].AdvanceInDirection(s.direction)
}

func (s *Snake) Grow() {
	s.body = append(s.body, Position{})
}

func (s *Snake) Colliding() bool {
	head := s.Head()

	for _, pos := range s.body[1:] {
		if head == pos {
			return true
		}
	}

	return false
}

func (p *Position) AdvanceInDirection(direction Direction) {
	switch direction {
	case Right:
		p.X++
	case Left:
		p.X--
	case Up:
		p.Y--
	case Down:
		p.Y++
	}
}

type Game struct {
	screen *Screen
	input  chan byte
	snake  Snake
	food   Food
}

func NewGame(s *Screen, r io.Reader) *Game {

	width, height := s.Size()

	initial := Position{
		height/2 + 1,
		width/2 + 1,
	}
	random := randomPosition(s)

	input := make(chan byte)
	go read(r, input)

	game := &Game{
		screen: s,
		input:  input,
		snake:  *NewSnake(initial),
		food:   Food(random),
	}

	return game
}

func read(r io.Reader, ch chan byte) {
	buf := make([]byte, 1)
	for {
		n, err := r.Read(buf)
		if err != nil || n != 1 {
			continue
		}

		ch <- buf[0]
	}
}

func randomPosition(s *Screen) Position {
	width, height := s.Size()
	random := Position{
		rand.Intn(height-2) + 2,
		rand.Intn(width-2) + 2,
	}
	return random
}

func (g *Game) Start() {
	g.screen.HideCursor()
	defer g.screen.ShowCursor()

	g.screen.EnableAlternativeBuffer()
	defer g.screen.DisableAlternativeBuffer()

	g.Loop()
}

func (g *Game) Loop() {

	framerate := 15
	desired := time.Second / time.Duration(framerate)

	timer := time.Now()

	for !g.GameOver() {
		g.Display()
		g.Update()

		elapsed := time.Since(timer)
		if elapsed < desired {
			time.Sleep(desired - elapsed)
		}
		timer = time.Now()
	}
}

func (g *Game) Update() {
	select {
	case input := <-g.input:
		g.HandleInput(input)
	default:
	}

	if g.snake.Head() == Position(g.food) {
		g.snake.Grow()

		g.RepositionFood()
	}

	g.snake.Advance()
}

func (g *Game) HandleInput(input byte) {
	switch input {
	case 'w':
		if g.snake.direction != Down {
			g.snake.direction = Up
		}
	case 's':
		if g.snake.direction != Up {
			g.snake.direction = Down
		}
	case 'a':
		if g.snake.direction != Right {
			g.snake.direction = Left
		}
	case 'd':
		if g.snake.direction != Left {
			g.snake.direction = Right
		}
	}
}

func (g *Game) RepositionFood() {
	g.food = Food(randomPosition(g.screen))
}

func (g *Game) GameOver() bool {
	width, height := g.screen.Size()

	head := g.snake.Head()

	if head.X <= 1 || head.X >= width {
		return true
	}
	if head.Y <= 1 || head.Y >= height {
		return true
	}

	if g.snake.Colliding() {
		return true
	}

	return false
}

func (g *Game) Display() {
	g.screen.Clear()

	g.displayBox()
	g.displayScore()
	g.displaySnake()
	g.displayFood()

	g.screen.Flush()
}

func (g *Game) displayFood() {
	g.screen.MoveTo(Position(g.food))
	g.screen.Printf("x")
}

func (g *Game) displaySnake() {
	for i, pos := range g.snake.body {
		g.screen.MoveTo(pos)

		if i == 0 {
			g.screen.Printf("O")
		} else {
			g.screen.Printf("o")
		}
	}
}

func (g *Game) displayBox() {
	width, height := g.screen.Size()
	g.screen.MoveTo(Position{0, 0})
	for i := 0; i < width; i++ {
		g.screen.Printf("-")
	}
	g.screen.MoveTo(Position{height, 0})
	for i := 0; i < width; i++ {
		g.screen.Printf("-")
	}
	for i := 1; i <= height; i++ {
		g.screen.MoveTo(Position{i, 0})
		g.screen.Printf("|")
		g.screen.MoveTo(Position{i, width})
		g.screen.Printf("|")
	}
}

func (g *Game) displayScore() {
	width, _ := g.screen.Size()
	g.screen.MoveTo(Position{1, width/2 - 3})
	g.screen.Printf("SCORE: %v", len(g.snake.body))
}
