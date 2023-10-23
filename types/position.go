package types

type Position struct {
	X, Y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p *Position) Advance(direction Direction) {
	switch direction {
	case Right:
		p.X++
	case Left:
		p.X--
	case Up:
		p.Y--
	case Down:
		p.Y++
	default:
		panic("invalid direction")
	}
}
