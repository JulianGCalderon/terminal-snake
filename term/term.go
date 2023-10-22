package term

import (
	"bufio"
	"fmt"
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

func MoveTo(t Terminal, x, y int) error {
	return PrintfCode(t, "[%v;%vH", y, x)
}

func Clear(t Terminal) error {
	defer t.Flush()
	return PrintfCode(t, "[2J")
}

func EnableAlternativeBuffer(t Terminal) error {
	defer t.Flush()
	return PrintfCode(t, "[?1049h")
}

func DisableAlternativeBuffer(t Terminal) error {
	defer t.Flush()
	return PrintfCode(t, "[?1049l")
}

func HideCursor(t Terminal) error {
	defer t.Flush()
	return PrintfCode(t, "[?25l")
}

func ShowCursor(t Terminal) error {
	defer t.Flush()
	return PrintfCode(t, "[?25h")
}

func PrintfCode(t Terminal, format string, a ...any) error {
	return Printf(t, ESC+format, a...)
}

func Printf(t Terminal, format string, a ...any) error {
	_, err := fmt.Fprintf(t, format, a...)
	return err
}
