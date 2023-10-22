package main

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	*bufio.Writer
	size [2]int
	fd   int
}

func (t *Terminal) Size() (int, int) {
	return t.size[0], t.size[1]
}

func NewTerminal(raw *os.File) (terminal *Terminal, err error) {
	terminal = &Terminal{}

	terminal.fd = int(raw.Fd())

	width, height, err := term.GetSize(terminal.fd)
	if err != nil {
		return
	}
	terminal.size[0] = width
	terminal.size[1] = height

	terminal.Writer = bufio.NewWriter(raw)

	return
}
