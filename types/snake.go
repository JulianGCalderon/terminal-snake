package types

type Snake struct {
	direction Direction
	body      []Position
}

func NewSnake(position Position) Snake {
	return Snake{
		body: []Position{position},
	}
}

func (s *Snake) Head() Position {
	return s.body[0]
}

func (s *Snake) Body() []Position {
	return s.body[1:]
}

func (s *Snake) Length() int {
	return len(s.body)
}

func (s *Snake) Direction() Direction {
	return s.direction
}

func (s *Snake) SetDirection(d Direction) {
	s.direction = d
}

func (s *Snake) Advance() {
	for i := len(s.body) - 1; i > 0; i-- {
		s.body[i] = s.body[i-1]
	}

	s.body[0].Advance(s.direction)
}

func (s *Snake) Grow() {
	s.body = append(s.body, Position{})
}

func (s *Snake) SelfColliding() bool {
	head := s.Head()

	for _, pos := range s.body[1:] {
		if head == pos {
			return true
		}
	}

	return false
}

func (s *Snake) OffBounds(b Bounds) bool {
	return !b.Contains(s.Head())
}
