package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
)

type movement struct {
	direction string
	amount    int
}

func (m movement) parse(line string) movement {
	parts := strings.Split(line, " ")
	m1 := movement{}
	m1.direction = parts[0]
	m1.amount, _ = strconv.Atoi(parts[1])
	return m1
}

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day02"
	input := aoc.GetStringArray("input.txt")
	movements := processInput(input)

	a = solveA(movements)
	b = solveB(movements)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solveA(ms []movement) int {
	var x, depth = 0, 0
	for _, m := range ms {
		switch m.direction {
		case "forward":
			x += m.amount
		case "down":
			depth += m.amount
		case "up":
			depth -= m.amount
		}
	}
	return x * depth
}

func solveB(ms []movement) int {

	var x, depth, aim = 0, 0, 0
	for _, m := range ms {
		switch m.direction {
		case "forward":
			x += m.amount
			depth += (aim * m.amount)
		case "down":
			aim += m.amount
		case "up":
			aim -= m.amount

		}
	}
	return x * depth
}

func processInput(lines []string) []movement {
	var m movement
	ms := []movement{}
	for _, line := range lines {
		m := m.parse(line)
		ms = append(ms, m)
	}
	return ms
}
