package term

import (
	"bufio"
	"io"
	"os"

	"golang.org/x/term"
)

const (
	ESC = "\x1b"
)

type Terminal interface {
	io.ReadWriter
	io.RuneReader
	Size() (width, height int, err error)
	Flush() error
	Restore() error
}

type terminal struct {
	*bufio.ReadWriter
	outputFd, inputFd int
	initial           *term.State
}

func (t *terminal) Size() (width, height int, err error) {
	return term.GetSize(t.outputFd)
}

func (t *terminal) Restore() error {
	return term.Restore(t.outputFd, t.initial)
}

func New() (Terminal, error) {
	t := &terminal{}

	input, output := os.Stdin, os.Stdout
	t.outputFd = int(os.Stdout.Fd())
	t.inputFd = int(os.Stdin.Fd())

	initial, err := term.MakeRaw(t.inputFd)
	if err != nil {
		return nil, err
	}
	t.initial = initial

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)
	t.ReadWriter = bufio.NewReadWriter(reader, writer)

	return t, nil
}
