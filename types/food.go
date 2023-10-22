package types

type Food Position

func NewFood(p Position) Food {
	return Food(p)
}

func (f Food) Position() Position {
	return Position(f)
}
