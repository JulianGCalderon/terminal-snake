package main

import (
	"fmt"
	"termsnake/term"
)

const ESC = "\x1b"

type Display struct {
	term.Terminal
}

func (d *Display) Size() (width, height int, err error) {
	width, height, err = d.Terminal.Size()
	if err != nil {
		return
	}

	width = width / 2

	return
}

func (d *Display) MoveTo(x, y int) error {
	x = (x-1)*2 + 1
	return d.PrintfCode("[%v;%vH", y, x)
}

func (d *Display) Clear() error {
	defer d.Flush()
	return d.PrintfCode("[2J")
}

func (d *Display) EnableAlternativeBuffer() error {
	defer d.Flush()
	return d.PrintfCode("[?1049h")
}

func (d *Display) DisableAlternativeBuffer() error {
	defer d.Flush()
	return d.PrintfCode("[?1049l")
}

func (d *Display) HideCursor() error {
	defer d.Flush()
	return d.PrintfCode("[?25l")
}

func (t *Display) ShowCursor() error {
	defer t.Flush()
	return t.PrintfCode("[?25h")
}

func (t *Display) PrintfCode(format string, a ...any) error {
	return t.Printf(ESC+format, a...)
}

func (t Display) Printf(format string, a ...any) error {
	_, err := fmt.Fprintf(t, format, a...)
	return err
}
