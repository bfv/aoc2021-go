package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bfv/aoc2021-go/aoc"
)

const boardsize = 5

type board struct {
	// x, y array
	numbers [5][5]int
	y       int
	bingo   bool
}

func (b *board) fillRow(line string) {
	line = strings.Trim(line, " ")
	line = strings.ReplaceAll(line, "  ", " ")
	strs := strings.Split(line, " ")
	for x, v := range strs {
		n, _ := strconv.Atoi(v)
		b.numbers[x][b.y] = n
	}
	b.y++
}

func (b *board) draw(n int) (board, bool) {
	for x := 0; x < boardsize; x++ {
		for y := 0; y < boardsize; y++ {
			if b.numbers[x][y] == n {
				b.numbers[x][y] = -1
			}
		}
	}
	b.bingo = b.checkBingo()
	return *b, b.bingo
}

func (b *board) checkBingo() bool { // return true if bingo

	for x := 0; x < boardsize; x++ {
		col := true
		row := true
		for y := 0; y < boardsize; y++ {
			if b.numbers[x][y] != -1 {
				col = false
			}
			if b.numbers[y][x] != -1 {
				row = false
			}
		}
		if col || row {
			return true
		}
	}
	return false
}

func (b *board) checkUnmarkedNumbers() int {
	var result int
	for x := 0; x < boardsize; x++ {
		for y := 0; y < boardsize; y++ {
			if b.numbers[x][y] != -1 {
				result += b.numbers[x][y]
			}
		}
	}
	return result
}

// display the bingo card
func (b board) Print() {
	for x := 0; x < boardsize; x++ {
		s := ""
		for y := 0; y < boardsize; y++ {
			s += fmt.Sprintf("%3d", b.numbers[x][y])
		}
		println(s)
	}
}

var numbers []int
var boards []board

func main() {

	start := time.Now()

	var a, b int = 0, 0

	day := "day04"
	input := aoc.GetStringArray("input.txt")

	numbers, boards = parseInput(input)
	a, b = executeDraw()

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func executeDraw() (int, int) {

	first := -1
	lastBoard, lastNumber := -1, -1
	for _, v := range numbers {
		bingoBoard := drawNumber(v)
		if bingoBoard > -1 {
			if first == -1 {
				first = v * boards[bingoBoard].checkUnmarkedNumbers()
			}
			if boardsLeft() == 0 {
				lastBoard = bingoBoard
				lastNumber = v
				break
			}
		}

	}

	last := boards[lastBoard].checkUnmarkedNumbers() * lastNumber
	return first, last
}

func boardsLeft() int {
	count := 0
	for _, b := range boards {
		if !b.bingo {
			count++
		}
	}
	return count
}

func drawNumber(n int) int {

	bingoBoard := -1
	for i, b := range boards {

		if b.bingo {
			continue
		}

		brd, bingo := b.draw(n)
		boards[i] = brd

		if bingo {
			bingoBoard = i
		}
	}
	return bingoBoard
}

func parseInput(input []string) ([]int, []board) {
	var boards []board

	numbers := aoc.GetNumberArray(input[0])

	brd := board{}
	for i, v := range input {

		if i < 2 {
			continue
		}

		if v == "" {
			boards = append(boards, brd)
			brd = board{}
		} else {
			brd.fillRow(v)
		}

	}
	boards = append(boards, brd)

	return numbers, boards
}
