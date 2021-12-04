package main

import (
	"fmt"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
)

var bitCount int

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day03"
	inputFile := "input.txt"
	bitCount = aoc.GetInputLineLength(inputFile)
	ints := aoc.GetIntArrayFromBinary(inputFile)

	count := getCount(ints)

	gamma := 0
	for i := 0; i < bitCount; i++ {
		if count[i] > len(ints)/2 {
			gamma |= 1 << (bitCount - 1 - i)
		}
	}

	epsilon := ((1 << bitCount) - 1) ^ gamma
	a = gamma * epsilon

	// part b

	intsMost := aoc.GetIntArrayFromBinary(inputFile)
	oxyRating := 0
	for len(intsMost) > 1 {

		for i := 0; i < bitCount && len(intsMost) > 1; i++ {

			count = getCount(intsMost)

			var tmp []int
			var half float32 = float32(len(intsMost)) / 2
			mostCommon := 0
			if float32(count[i]) >= half {
				mostCommon = 1
			}

			for _, x := range intsMost {
				// bx := strconv.FormatInt(int64(x), 2)
				// bx = bx
				masked := (1 << (bitCount - 1 - i)) & x
				if (masked == 0 && mostCommon == 0) || (masked > 0 && mostCommon == 1) {
					tmp = append(tmp, x)
				}
			}
			intsMost = tmp
		}
		oxyRating = intsMost[0]
	}

	intsLeast := aoc.GetIntArrayFromBinary(inputFile)
	co2Rating := 0
	for len(intsLeast) > 1 {

		for i := 0; i < bitCount && len(intsLeast) > 1; i++ {

			count = getCount(intsLeast)

			var tmp []int
			var half float32 = float32(len(intsLeast)) / 2
			leastCommon := 1
			if float32(count[i]) >= half {
				leastCommon = 0
			}

			for _, x := range intsLeast {
				// bx := strconv.FormatInt(int64(x), 2)
				// bx = bx
				mask := (1 << (bitCount - 1 - i))
				masked := mask & x
				if (masked == 0 && leastCommon == 0) || (masked > 0 && leastCommon == 1) {
					tmp = append(tmp, x)
				}
			}
			intsLeast = tmp
		}
		co2Rating = intsLeast[0]
	}

	b = oxyRating * co2Rating

	fmt.Printf("oxy rating: %v\n", oxyRating)
	fmt.Printf("co2 scrubber rating: %v\n", co2Rating)

	elapsed := time.Since(start)

	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func getCount(ints []int) []int {
	count := make([]int, bitCount)
	for _, x := range ints {
		for i := 0; i < bitCount; i++ {
			if x&(1<<(bitCount-1-i)) > 0 {
				count[i]++
			}
		}
	}
	return count
}

// func getMost(ints []int) int {
// 	most := 0
// 	count := getCount(ints)
// 	for i := 0; i < bitCount; i++ {
// 		if count[i] > len(ints)/2 {
// 			most |= 1 << (bitCount - 1 - i)
// 		}
// 	}
// 	fmt.Println(strconv.FormatInt(int64(most), 2))
// 	return most
// }
