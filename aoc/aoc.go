package aoc

import (
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

func GetStringArray(filename string) []string {
	var sep string

	if runtime.GOOS == "windows" {
		sep = "\r\n"
	} else {
		sep = "\n"
	}

	buf, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(buf), sep)

	return input
}

func GetIntArray(filename string) []int {
	ints := []int{}
	lines := GetStringArray(filename)
	for _, line := range lines {
		x, _ := strconv.Atoi(line)
		ints = append(ints, x)
	}
	return ints
}
