package static

type Stone struct {
	x int
	y int
}

func NewStone(x int, y int) *Stone {
	return &Stone{
		x: x,
		y: y,
	}
}

func (s *Stone) Coordinates() (int, int) {
	return s.x, s.y
}
