package main

import (
	"fmt"
)

const esc = "\x1b"

type Screen struct {
	*Terminal
}

func NewScreen(terminal *Terminal) *Screen {
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
	s.Printf(esc+"[%v;%vH", position.Y, position.X)
}

func (s *Screen) Clear() {
	s.Printf(esc + "[2J")
}

func (s *Screen) Reset() {
	s.MoveTo(Position{0, 0})
	s.Clear()
	s.Flush()
}
