package types

import "fmt"

type Direction int8

const (
	Right Direction = iota
	Up    Direction = iota
	Down  Direction = iota
	Left  Direction = iota
)

func NewDirection(r rune) (new Direction, err error) {
	switch r {
	case 'a':
		new = Left
	case 's':
		new = Down
	case 'd':
		new = Right
	case 'w':
		new = Up
	default:
		err = fmt.Errorf("unknown direction")
	}

	return
}

func (d Direction) Opposite() Direction {

	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}

	return d
}
