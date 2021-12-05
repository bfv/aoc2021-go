package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
	"github.com/bfv/aoc2021-go/lib"
	"github.com/bfv/aoc2021-go/shapes"
)

type seafloor shapes.Grid

func (s seafloor) Print() {
	g := shapes.Grid(s)
	g.Print()
}

var seafloorA, seafloorB seafloor
var lines []line

type line struct {
	x0, y0   int
	x1, y1   int
	horOrVer bool
}

func (l line) getLine(s string) line {
	parts := strings.Split(s, "->")
	parts0 := strings.Split(parts[0], ",")
	parts1 := strings.Split(parts[1], ",")
	l0 := line{}
	l0.x0 = getInt(parts0[0])
	l0.y0 = getInt(parts0[1])
	l0.x1 = getInt(parts1[0])
	l0.y1 = getInt(parts1[1])
	l0.horOrVer = l0.x0 == l0.x1 || l0.y0 == l0.y1
	return l0
}

func (l line) apply(f *seafloor) {

	deltaX := l.x1 - l.x0
	deltaY := l.y1 - l.y0

	dx := getStepDelta(deltaX)
	dy := getStepDelta(deltaY)

	dist := lib.Max(lib.Abs(deltaX), lib.Abs(deltaY))
	x, y := l.x0, l.y0
	for i := 0; i <= dist; i++ {
		(*f)[co(x, y)]++
		x += dx
		y += dy
	}

}

func getStepDelta(n int) int {
	delta := 0
	if n != 0 {
		if n > 0 {
			delta = 1
		} else {
			delta = -1
		}
	}
	return delta
}

func co(x int, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func main() {

	start := time.Now()

	var a, b int = 0, 0
	var l line
	day := "day05"

	input := aoc.GetStringArray("input.txt")
	lines = []line{}
	seafloorA, seafloorB = seafloor{}, seafloor{}

	for _, s := range input {
		l = l.getLine(s)
		lines = append(lines, l)
	}

	a = solveA()
	b = solveB()

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solveA() int {

	for _, l := range lines {
		if l.horOrVer {
			l.apply(&seafloorA)
		}
	}

	danger := 0
	for _, v := range seafloorA {
		if v > 1 {
			danger++
		}
	}

	return danger
}

func solveB() int {

	for _, l := range lines {
		l.apply(&seafloorB)
	}

	danger := 0
	for _, v := range seafloorB {
		if v > 1 {
			danger++
		}
	}

	return danger
}

func getInt(s string) int {
	s = strings.Trim(s, " ")
	i, _ := strconv.Atoi(s)
	return i
}
