package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	// implement me
	s.start.x = s.start.x + int(s.a)
	s.start.y = s.start.y + int(s.a)

	return Point{x: s.start.x, y: s.start.y}
}

func (s Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	// implement
	return 4 * s.a
}
