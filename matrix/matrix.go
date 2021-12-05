package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	data [][]int
	rows int
	cols int
}

func (m Matrix) New(rows int, cols int) Matrix {
	mr := Matrix{}
	mr.rows, mr.cols = rows, cols
	mr.data = make([][]int, rows)
	for row := 0; row < rows; row++ {
		mr.data[row] = make([]int, cols)
	}
	return mr
}

func (m Matrix) Print() {
	for row := 0; row < m.rows; row++ {
		s := "|"
		for col := 0; col < m.cols; col++ {
			s += fmt.Sprintf("%3d", m.data[row][col])
		}
		s += " |"
		fmt.Println(s)
	}
	fmt.Println()
}

func (m1 Matrix) Multiply(m2 Matrix) Matrix {
	if m1.cols != m2.rows {
		panic(fmt.Sprintf("cannot multiply a %vx%v by a %vx%v matrices", m1.rows, m1.cols, m2.rows, m2.cols))
	}

	mr := m1.New(m1.rows, m2.cols)

	for row := 0; row < mr.rows; row++ {
		for col := 0; col < mr.cols; col++ {
			r := 0
			for i := 0; i < m1.cols; i++ {
				r += m1.data[row][i] * m2.data[i][col]
			}
			mr.data[row][col] = r
		}
	}
	return mr
}

func (m Matrix) Transpose() Matrix {
	mt := m.New(m.cols, m.rows)
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			mt.data[col][row] = m.data[row][col]
		}
	}
	return mt
}

func (m *Matrix) FillRowWithString(s string, row int, sep string) {
	s = strings.ReplaceAll(strings.Trim(s, " "), "  ", " ") // remove leading and double spaces
	strs := strings.Split(s, sep)
	for col, v := range strs {
		n, _ := strconv.Atoi(strings.Trim(v, " "))
		m.data[row][col] = n
	}
}
