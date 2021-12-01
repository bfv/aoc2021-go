package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
)

func main() {

	start := time.Now()

	var a, b int = -1, -1

	day := "day01"
	buf, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(buf), "\r\n")

	prev, prevB := -1, -1

	ints := aoc.GetIntArray(input)
	for i, x := range ints {
		// part a
		if x > prev {
			a++
		}
		prev = x

		// part b
		if i > 1 {
			y := ints[i-2] + ints[i-1] + ints[i]
			if y > prevB {
				b++
			}
			prevB = y
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}
