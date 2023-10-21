package main

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	*bufio.Writer
	width  int
	height int
	fd     int
	state  *term.State
}

func (t *Terminal) Size() (int, int) {
	return t.width, t.height
}

func NewTerminal(raw *os.File) (terminal *Terminal, err error) {
	terminal = &Terminal{}

	terminal.fd = int(raw.Fd())

	width, height, err := term.GetSize(terminal.fd)
	if err != nil {
		return
	}
	terminal.width = width
	terminal.height = height

	terminal.Writer = bufio.NewWriter(raw)

	return
}

func (t *Terminal) Save() error {
	state, err := term.GetState(int(t.fd))
	t.state = state

	return err
}

func (t *Terminal) Restore() error {
	return term.Restore(int(t.fd), t.state)
}
