package main

import (
	"log"
	"os"

	"golang.org/x/term"
)

func main() {
	terminal, err := NewTerminal(os.Stdout)
	if err != nil {
		log.Fatalf("NewTerminal: %v\n", err)
	}
	screen := NewScreen(terminal)

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("MakeRaw: %v\n", err)
	}

	game := NewGame(screen, os.Stdin)
	game.Start()

	err = term.Restore(int(os.Stdin.Fd()), state)
	if err != nil {
		log.Fatalf("Restore: %v\n", err)
	}
}
