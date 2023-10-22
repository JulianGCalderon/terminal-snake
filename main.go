package main

import (
	"log"
	"termsnake/term"
)

func main() {
	terminal, err := term.New()
	if err != nil {
		log.Fatalf("Could not create terminal: %v", err)
	}
	defer terminal.Restore()

	game, err := NewGame(terminal)
	if err != nil {
		log.Fatalf("Could not create game: %v", err)
	}

	game.Start()
}
