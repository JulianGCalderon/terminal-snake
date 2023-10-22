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

func (s *Screen) MoveTo(position Position) {
	s.Printf(esc+"[%v;%vH", position.Y, position.X)
}

func (s *Screen) Clear() {
	s.Printf(esc + "[2J")
	s.Flush()
}

func (s *Screen) HideCursor() {
	s.Printf(esc + "[?25l")
	s.Flush()
}

func (s *Screen) ShowCursor() {
	s.Printf(esc + "[?25h")
	s.Flush()
}

func (s *Screen) EnableAlternativeBuffer() {
	s.Printf(esc + "[?1049h")
	s.Flush()
}

func (s *Screen) DisableAlternativeBuffer() {
	s.Printf(esc + "[?1049l")
	s.Flush()
}
