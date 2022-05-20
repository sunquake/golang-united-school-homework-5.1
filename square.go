package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (r Square) End() Point {
	return Point{x:r.start.x+r.start.a, y:r.start.y+r.start.a}
}

func (r Square) Area() uint {
	return r.start.a * r.start.a
}

func (r Square) Perimeter() uint {
	return r.start.a * 4
}
