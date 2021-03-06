package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
	"github.com/bfv/aoc2021-go/lib"
)

var fuelCostsB map[int]int

func main() {
	start := time.Now()

	var a, b int = 0, 0

	day := "day07"
	input := aoc.GetIntArrayFromStringArray(strings.Split(aoc.GetStringArray("input.txt")[0], ","))
	fuelCostsB = map[int]int{}

	a, b = solve(input)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solve(input []int) (int, int) {
	highest := aoc.GetHighestIntFromArray(input)
	lowestFuelA, lowestFuelB := math.MaxInt, math.MaxInt
	for a := 0; a <= highest; a++ {
		fuelA, fuelB := 0, 0
		for _, v := range input {
			d := lib.Abs(v - a)
			fuelA += d
			fuelB += d * (d + 1) / 2
		}
		lowestFuelA = lib.Min(fuelA, lowestFuelA)
		lowestFuelB = lib.Min(fuelB, lowestFuelB)
	}
	return lowestFuelA, lowestFuelB
}
