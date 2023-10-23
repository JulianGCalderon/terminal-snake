package main

import (
	"fmt"
	"io"
	"termsnake/types"
	"time"

	"termsnake/term"
)

const FrameRate = 15

type Game struct {
	terminal      term.Terminal
	input         chan rune
	width, height int
	snake         types.Snake
	food          types.Food
}

func NewGame(t term.Terminal) (game *Game, err error) {
	game = &Game{terminal: t}

	err = game.updateSize()
	if err != nil {
		return
	}

	initial := game.bounds().Middle()
	game.snake = types.NewSnake(initial)

	game.SpawnFood()

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

func (g *Game) bounds() types.Bounds {
	return types.Bounds{
		MinX: 1,
		MaxX: g.width / 2,
		MinY: 1,
		MaxY: g.height,
	}
}

func (g *Game) Start() {

	term.EnableAlternativeBuffer(g.terminal)

	term.HideCursor(g.terminal)

	g.Loop()

	term.DisableAlternativeBuffer(g.terminal)
	term.ShowCursor(g.terminal)

	fmt.Printf("Game Over! - Score: %v", g.score())
}

func (g *Game) Loop() {

	ticker := time.NewTicker(time.Second / FrameRate)

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

	g.updateSize()

	if g.snake.Head() == g.food.Position() {
		g.snake.Grow()
		g.SpawnFood()
	}

	g.snake.Advance()
}

func (g *Game) updateSize() error {
	width, height, err := g.terminal.Size()
	if err != nil {
		return err
	}

	g.width, g.height = width, height

	return nil
}

func (g *Game) SpawnFood() {
	g.food = types.NewFood(g.bounds().Random())
}

func (g *Game) HandleInput(input rune) {
	direction, err := types.NewDirection(input)
	if err != nil {
		return
	}

	if g.snake.Direction.Opposite() != direction {
		g.snake.Direction = direction
	}
}

func (g *Game) GameOver() bool {
	return g.snake.SelfColliding() || g.snake.OffBounds(g.bounds())
}

func (g *Game) Display() {
	term.Clear(g.terminal)

	g.displayBounds()
	g.displayFood()
	g.displaySnake()
	g.displayScore()

	g.terminal.Flush()
}

func (g *Game) displayFood() {
	term.MoveTo(g.terminal, g.food.X, g.food.Y)
	term.Printf(g.terminal, "⚪")
}

func (g *Game) displaySnake() {
	term.MoveTo(g.terminal, g.snake.Head().X, g.snake.Head().Y)
	term.Printf(g.terminal, "██")

	for _, pos := range g.snake.Body() {
		term.MoveTo(g.terminal, pos.X, pos.Y)
		term.Printf(g.terminal, "▒▒")
	}
}

func (g *Game) displayBounds() {
	bounds := g.bounds()

	term.MoveTo(g.terminal, bounds.MinX, bounds.MinY)

	term.Printf(g.terminal, "┌─")
	for i := 0; i < bounds.Width(); i++ {
		term.Printf(g.terminal, "──")
	}
	term.Printf(g.terminal, "─┐")

	term.MoveTo(g.terminal, bounds.MinX, bounds.MaxY)

	term.Printf(g.terminal, "└─")
	for i := 0; i < bounds.Width(); i++ {
		term.Printf(g.terminal, "──")
	}
	term.Printf(g.terminal, "─┘")

	for i := 0; i < bounds.Height(); i++ {
		term.MoveTo(g.terminal, 1, i+2)
		term.Printf(g.terminal, "│")
		term.MoveTo(g.terminal, bounds.MaxX, i+2)
		term.Printf(g.terminal, " │")
	}
}

func (g *Game) displayScore() {
	term.MoveTo(g.terminal, g.bounds().Middle().X, 1)
	term.Printf(g.terminal, " SCORE: %v ", g.score())
}

func (g *Game) score() int {
	return g.snake.Length() - 1
}
