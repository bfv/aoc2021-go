package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
)

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day01"
	ints := aoc.GetIntArray("input.txt")
	for i, _ := range ints {

		if i > 0 && ints[i] > ints[i-1] {
			a++
		}

		if i > 2 && ints[i] > ints[i-3] { // solution authored by Roel
			b++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}
