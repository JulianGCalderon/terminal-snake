package main

import (
	"io"
	"math/rand"
	"termsnake/types"
	"time"

	"termsnake/term"
)

type Game struct {
	terminal      term.Terminal
	input         chan rune
	width, height int
	snake         types.Snake
	food          types.Food
}

func NewGame(t term.Terminal) (game *Game, err error) {
	game = &Game{}

	game.terminal = t

	width, height, err := t.Size()
	if err != nil {
		return
	}
	game.width, game.height = width, height

	initial := game.middlePosition()
	game.snake = types.NewSnake(initial)

	game.spawnFood()

	game.input = make(chan rune)
	go read(t, game.input)

	return
}

func read(r io.RuneReader, ch chan rune) {
	for {
		r, _, err := r.ReadRune()
		if err != nil {
			continue
		}

		ch <- r
	}
}

func (g *Game) middlePosition() types.Position {
	x := g.width / 2
	y := g.height / 2

	return types.Position{X: x, Y: y}
}

func (g *Game) randomPosition() types.Position {
	x := rand.Intn(g.width-1) + 2
	y := rand.Intn(g.height-1) + 2

	return types.Position{X: x, Y: y}
}

func (g *Game) Start() {
	term.HideCursor(g.terminal)
	defer term.ShowCursor(g.terminal)

	term.EnableAlternativeBuffer(g.terminal)
	defer term.DisableAlternativeBuffer(g.terminal)

	g.Loop()
}

func (g *Game) Loop() {

	ticker := time.NewTicker(50 * time.Millisecond)

	for !g.GameOver() {
		g.Display()
		g.Update()
		<-ticker.C
	}
}

func (g *Game) Update() {

	select {
	case input := <-g.input:
		g.HandleInput(input)
	default:
	}

	if g.snake.Head() == g.food.Position() {
		g.snake.Grow()
		g.spawnFood()
	}

	g.snake.Advance()
}

func (g *Game) HandleInput(input rune) {
	direction, err := types.NewDirection(input)
	if err != nil {
		return
	}

	if g.snake.Direction().Opposite() != direction {
		g.snake.SetDirection(direction)
	}
}

func (g *Game) spawnFood() {
	random := g.randomPosition()
	g.food = types.NewFood(random)
}

func (g *Game) GameOver() bool {
	return g.snake.SelfColliding()
}

func (g *Game) Display() {
	term.Clear(g.terminal)

	g.displayFood()
	g.displaySnake()

	g.terminal.Flush()
}

func (g *Game) displayFood() {
	term.MoveTo(g.terminal, g.food.X, g.food.Y)
	term.Printf(g.terminal, "*")
}

func (g *Game) displaySnake() {
	term.MoveTo(g.terminal, g.snake.Head().X, g.snake.Head().Y)
	term.Printf(g.terminal, "O")

	for _, pos := range g.snake.Body() {
		term.MoveTo(g.terminal, pos.X, pos.Y)
		term.Printf(g.terminal, "o")
	}
}

// func (g *Game) displayBox() {
// 	width, height := g.screen.Size()
// 	g.screen.MoveTo(Position{0, 0})
// 	for i := 0; i < width; i++ {
// 		g.screen.Printf("-")
// 	}
// 	g.screen.MoveTo(Position{height, 0})
// 	for i := 0; i < width; i++ {
// 		g.screen.Printf("-")
// 	}
// 	for i := 1; i <= height; i++ {
// 		g.screen.MoveTo(Position{i, 0})
// 		g.screen.Printf("|")
// 		g.screen.MoveTo(Position{i, width})
// 		g.screen.Printf("|")
// 	}
// }

// func (g *Game) displayScore() {
// 	width, _ := g.screen.Size()
// 	g.screen.MoveTo(Position{1, width/2 - 3})
// 	g.screen.Printf("SCORE: %v", len(g.snake.body))
// }
