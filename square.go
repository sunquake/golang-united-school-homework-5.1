package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (r Square) End() Point {
	return Point{x:r.start.x+r.a, y:r.start.y+r.a}
}

func (r Square) Area() uint {
	return r.a * r.a
}

func (r Square) Perimeter() uint {
	return r.a * 4
}
