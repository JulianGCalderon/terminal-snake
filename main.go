package main

import (
	"os"
	"time"
)

func main() {
	terminal := NewTerminal(os.Stdout)
	screen := NewScreen(terminal)

	screen.Clear()
	screen.HideCursor()
	screen.MoveTo(Position{10, 10})
	screen.Printf("¡Hola Mundo!")

	time.Sleep(5 * time.Second)
}
