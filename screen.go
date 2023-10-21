package main

import (
	"fmt"
	"io"
	"os"
)

const esc = "\x1b"

type Terminal interface {
	io.Writer
	// size() (int, int)
}

func NewTerminal(terminal *os.File) Terminal {
	return terminal
}

type Screen struct {
	Terminal
}

func NewScreen(terminal Terminal) *Screen {
	return &Screen{terminal}
}

func (s *Screen) Printf(format string, a ...any) {
	fmt.Fprintf(s.Terminal, format, a...)
}

func (s *Screen) HideCursor() {
	s.Printf(esc + "[?25l")
}

func (s *Screen) ShowCursor() {
	s.Printf(esc + "[?25h")
}

func (s *Screen) MoveTo(position Position) {
	s.Printf(esc+"[%v;%vH", position.X, position.X)
}

func (s *Screen) Clear() {
	s.Printf(esc + "[2J")
}
