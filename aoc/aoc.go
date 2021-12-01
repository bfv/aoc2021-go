package aoc

import "strconv"

func GetIntArray(lines []string) []int {
	ints := []int{}
	for _, line := range lines {
		x, _ := strconv.Atoi(line)
		ints = append(ints, x)
	}
	return ints
}
