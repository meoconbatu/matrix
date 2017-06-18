package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

type Matrix struct {
	m    [][]int
	x, y int
}

const seperatorRow, seperatorCol = "\n", " "

func (m Matrix) Rows() [][]int {
	output := initial(m.x, m.y)
	for i, row := range m.m {
		for j, col := range row {
			output[i][j] = col
		}
	}
	return output
}
func (m Matrix) Cols() [][]int {
	output := initial(m.y, m.x)
	for i, row := range m.m {
		for j, col := range row {
			output[j][i] = col
		}
	}
	return output
}
func initial(x, y int) [][]int {
	output := make([][]int, x)
	for i := range output {
		output[i] = make([]int, y)
	}
	return output
}
func New(in string) (*Matrix, error) {
	output := new(Matrix)
	if in == "" {
		return output, nil
	}
	rows := strings.Split(in, seperatorRow)
	output.x = len(rows)
	output.m = make([][]int, output.x)
	for i, row := range rows {
		cols := strings.Split(strings.Trim(row, " "), seperatorCol)
		numCol := len(cols)
		if numCol == 0 {
			return nil, errors.New("row empty")
		}
		if output.y == 0 {
			output.y = numCol
		}
		if output.y != numCol {
			return nil, errors.New("uneven rows")
		}
		output.m[i] = make([]int, output.y)
		for j, col := range cols {
			a, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			output.m[i][j] = a
		}
	}
	return output, nil
}
func (m Matrix) Set(row, col, val int) bool {
	if row >= m.x || col >= m.y || row < 0 || col < 0 {
		return false
	}
	m.m[row][col] = val
	return true
}
