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
	input := aoc.GetIntArrayFromString(strings.Split(aoc.GetStringArray("input.txt")[0], ","))
	fuelCostsB = map[int]int{}

	a, b = solve(input)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solve(input []int) (int, int) {
	highest := aoc.GetHighestIntFromArray(input)
	lowestFuelA := math.MaxInt
	lowestFuelB := math.MaxInt
	for a := 0; a <= highest; a++ {
		fuelA := 0
		fuelB := 0
		for _, v := range input {
			fuelA += lib.Abs(v - a)
			fuelB += fuelBCost(lib.Abs(v - a))
		}
		if fuelA < lowestFuelA {
			lowestFuelA = fuelA
		}
		if fuelB < lowestFuelB {
			lowestFuelB = fuelB
		}
	}
	return lowestFuelA, lowestFuelB
}

func fuelBCost(d int) int {
	return d * (d + 1) / 2
}
