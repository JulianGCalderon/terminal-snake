package types

import "math/rand"

type Bounds struct {
	MinX, MaxX int
	MinY, MaxY int
}

func (b Bounds) Contains(p Position) bool {
	return p.X > b.MinX && p.Y > b.MinY &&
		p.X < b.MaxX && p.Y < b.MaxY
}

func (b Bounds) Middle() Position {
	return Position{
		X: (b.MaxX-b.MinX)/2 + b.MinX,
		Y: (b.MaxY-b.MinY)/2 + b.MinY,
	}
}

func (b Bounds) Random() Position {
	x := rand.Intn(b.Width()) + b.MinX + 1
	y := rand.Intn(b.Height()) + b.MinY + 1

	return NewPosition(x, y)
}

func (b Bounds) Width() int {
	return b.MaxX - b.MinX - 1
}

func (b Bounds) Height() int {
	return b.MaxY - b.MinY - 1
}
