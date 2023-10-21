package main

import (
	"log"
	"os"
	"time"
)

func main() {
	terminal, err := NewTerminal(os.Stdout)
	if err != nil {
		log.Fatalf("NewTerminal: %v\n", err)
	}

	terminal.Save()

	screen := NewScreen(terminal)

	game := NewGame(screen)

	game.Display()

	time.Sleep(1 * time.Second)

	screen.Reset()
}
