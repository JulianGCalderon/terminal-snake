package types

type Bounds struct {
	MinX, MaxX int
	MinY, MaxY int
}

func (b *Bounds) Contains(p Position) bool {
	return p.X >= b.MinX && p.Y >= b.MinY &&
		p.X <= b.MaxX && p.Y <= b.MaxY

}
