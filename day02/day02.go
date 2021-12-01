package aoc

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day01"
	buf, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(buf), "\r\n")

	
	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solve(input []int) (int, int) {
	answerA, answerB := -1, -1
	return answerA, answerB
}
