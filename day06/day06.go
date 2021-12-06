package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
	"github.com/bfv/aoc2021-go/lib"
)

type Fish struct {
	timer int
}

func (f *Fish) addDay() bool {
	createNew := false
	if f.timer == 0 {
		f.timer = 7
		createNew = true
	}
	f.timer--
	return createNew
}

var school []Fish
var timings []int

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day06"
	input := aoc.GetStringArray("input.txt")[0]
	inputStrings := strings.Split(input, ",")
	for _, t := range inputStrings {
		school = append(school, Fish{lib.Atoi(t)})
	}

	a = solveA()

	timings = make([]int, 9)
	for _, f := range inputStrings {
		timings[lib.Atoi(f)]++
	}
	b = solveB(256)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func solveA() int {

	answerA := 0
	newSchool := school
	for day := 1; day <= 80; day++ {
		tmp := newSchool
		for i, f := range tmp {
			createNew := f.addDay()
			tmp[i] = f
			newSchool[i] = f
			if createNew {
				newSchool = append(newSchool, Fish{8})
			}
		}
	}
	answerA = len(newSchool)
	return answerA
}

func solveB(days int) int {
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
