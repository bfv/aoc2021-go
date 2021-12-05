package shapes

import "fmt"

type Line struct {
	p0, p1 Point
}

func (l Line) IsHorOrVer() bool {
	return l.p0.x == l.p1.x || l.p0.y == l.p1.y
}

func (l Line) IsDiagonal() bool {
	dx := l.p0.x - l.p1.x
	dy := l.p0.y - l.p1.y
	return (dx == dy || dx == -dy)
}

func (l Line) Print() {
	fmt.Printf("%v -> %v\n", l.p0.Print(), l.p1.Print())
}
