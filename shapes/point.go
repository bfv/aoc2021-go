package shapes

import (
	"fmt"
	"math"

	"github.com/bfv/aoc2021-go/lib"
)

type Point struct {
	x, y int
}

func (p0 Point) Distance(p1 Point) float64 {
	dx, dy := p0.x-p1.x, p0.y-p1.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func (p0 Point) GridDistance(p1 Point) int {
	dx, dy := p0.x-p1.x, p0.y-p1.y
	return lib.Abs(dx) + lib.Abs(dy)
}

func (p Point) Print() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}
