package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
	"github.com/bfv/aoc2021-go/lib"
)

var timings []int

func main() {

	start := time.Now()

	day := "day06"
	inputStrings := strings.Split(aoc.GetStringArray("input.txt")[0], ",")

	timings = getTiming(inputStrings)
	a := solve(80)

	timings = getTiming(inputStrings)
	b := solve(256)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solve(days int) int {
	for day := 1; day <= days; day++ {
		shift()
	}
	return lib.CalcArraySum(timings)
}

func shift() {
	new := timings[0]
	for j := 0; j < 8; j++ {
		timings[j] = timings[j+1]
	}
	timings[6] += new
	timings[8] = new
}

func getTiming(inputStrings []string) []int {
	timings = make([]int, 9)
	for _, f := range inputStrings {
		timings[lib.Atoi(f)]++
	}
	return timings
}
