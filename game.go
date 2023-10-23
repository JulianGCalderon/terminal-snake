package main

import (
	"fmt"
	"io"
	"termsnake/types"
	"time"

	"termsnake/term"
)

const FrameRate = 20

type Game struct {
	display       Display
	input         chan rune
	width, height int
	snake         types.Snake
	food          types.Food
}

func NewGame(t term.Terminal) (game *Game, err error) {
	game = &Game{display: Display{t}}

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
		MaxX: g.width,
		MinY: 1,
		MaxY: g.height,
	}
}

func (g *Game) Start() {

	g.display.EnableAlternativeBuffer()
	g.display.HideCursor()

	g.Loop()

	g.display.DisableAlternativeBuffer()
	g.display.ShowCursor()

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
	width, height, err := g.display.Size()
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
	g.display.Clear()

	g.displayBounds()
	g.displayFood()
	g.displaySnake()
	g.displayScore()

	g.display.Flush()
}

func (g *Game) displayFood() {
	g.display.MoveTo(g.food.X, g.food.Y)
	g.display.Printf("⚪")
}

func (g *Game) displaySnake() {
	g.display.MoveTo(g.snake.Head().X, g.snake.Head().Y)
	g.display.Printf("██")

	for _, pos := range g.snake.Body() {
		g.display.MoveTo(pos.X, pos.Y)
		g.display.Printf("▒▒")
	}
}

func (g *Game) displayBounds() {
	bounds := g.bounds()

	g.display.MoveTo(bounds.MinX, bounds.MinY)

	g.display.Printf(" ┌")
	for i := 0; i < bounds.Width(); i++ {
		g.display.Printf("──")
	}
	g.display.Printf("┐ ")

	g.display.MoveTo(bounds.MinX, bounds.MaxY)

	g.display.Printf(" └")
	for i := 0; i < bounds.Width(); i++ {
		g.display.Printf("──")
	}
	g.display.Printf("┘ ")

	for i := 0; i < bounds.Height(); i++ {
		g.display.MoveTo(1, i+2)
		g.display.Printf(" │")
		g.display.MoveTo(bounds.MaxX, i+2)
		g.display.Printf("│ ")
	}
}

func (g *Game) displayScore() {
	g.display.MoveTo(g.bounds().Middle().X-1, 1)
	g.display.Printf(" SCORE: %v ", g.score())
}

func (g *Game) score() int {
	return g.snake.Length() - 1
}
