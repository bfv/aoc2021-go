package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
	"github.com/bfv/aoc2021-go/lib"
)

var incomplete []lib.Stack

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day10"
	input := aoc.GetStringArray("input.txt")

	// logic here
	a, b = solve(input)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solve(input []string) (int, int) {
	answerA, answerB := 0, 0

	// part a
	for _, line := range input {
		st := lib.Stack{}
		parts := strings.Split(line, "")
		corrupt := false
		for _, part := range parts {
			switch part {
			case "(", "[", "{", "<":
				st.Push(part)
			case ")", "]", "}", ">":
				popped := fmt.Sprintf("%v", st.Pop())
				val := checkValue(popped, part)
				answerA += val
				corrupt = (val > 0)
			}
			if corrupt {
				break
			}
		}
		if !corrupt {
			incomplete = append(incomplete, st)
		}
	}

	// part b
	scores := []int{}
	for _, st := range incomplete {
		score := 0
		for !st.IsEmpty() {
			val := getCloseValue(fmt.Sprintf("%v", st.Pop()))
			score = (5 * score) + val
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	answerB = scores[(len(scores)-1)/2]
	return answerA, answerB
}

func checkValue(open string, close string) int {
	val := 0
	posOpen := lib.IndexOfString(open, []string{"(", "[", "{", "<"})
	posClose := lib.IndexOfString(close, []string{")", "]", "}", ">"})

	if posOpen != posClose {
		val = []int{3, 57, 1197, 25137}[posClose]
	}
	return val
}

func getCloseValue(open string) int {
	return lib.IndexOfString(open, []string{"(", "[", "{", "<"}) + 1
}
