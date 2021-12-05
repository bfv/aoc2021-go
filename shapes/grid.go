package shapes

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bfv/aoc2021-go/lib"
)

type Grid map[string]int

func (g Grid) GetSize() (int, int) {
	pMin, pMax := g.GetCorners()
	return pMax.x - pMin.x, pMax.y - pMin.y
}

func (g Grid) GetCorners() (Point, Point) {

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for co := range g {
		x, y := Coxy(co)
		minX = lib.Min(x, minX)
		minY = lib.Min(y, minY)
		maxX = lib.Max(x, maxX)
		maxY = lib.Max(y, maxY)
	}
	return Point{minX, minY}, Point{maxX, maxY}
}

func (g Grid) Print() {
	p0, p1 := g.GetCorners()
	for y := lib.Min(0, p0.y); y <= p1.y; y++ {
		row := ""
		for x := lib.Min(0, p0.x); x <= p1.x; x++ {
			s := ""
			if v, ok := g[XYco(x, y)]; ok {
				s = strconv.Itoa(v)
			} else {
				s = "."
			}
			row += s
		}
		fmt.Println(row)
	}
}

func XYco(x int, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func Coxy(co string) (int, int) {
	parts := strings.Split(co, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}
