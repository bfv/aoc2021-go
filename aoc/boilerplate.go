package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {

	var a, b int = -1, -1

	day := "day01"
	buf, _ := ioutil.ReadFile("_input.txt")
	lines := strings.Split(string(buf), "\r\n")

	start := time.Now()
	for _, line := range lines {
		fmt.Println(line)
	}
	elapsed := time.Since(start)

	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solve(input []int) (int, int) {
	answerA, answerB := -1, -1
	return answerA, answerB
}
