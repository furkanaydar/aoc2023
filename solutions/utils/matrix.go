package utils

import (
	"fmt"
	"github.com/fatih/color"
	"hash/fnv"
)

type Cell struct {
	X   int
	Y   int
	Val rune
}

type Row []rune
type Column []rune
type Matrix []Row

func (row Row) ForEachChar(function func(c rune)) {
	for _, ch := range row {
		function(ch)
	}
}

func (row Row) ForEachCell(function func(i int, c rune)) {
	for index, ch := range row {
		function(index, ch)
	}
}

func (row Row) FindElement(elem rune) int {
	for i, cell := range row {
		if cell == elem {
			return i
		}
	}

	return -1
}

func (row Row) FindElementAll(elem rune) []int {
	var result []int

	for i, cell := range row {
		if cell == elem {
			result = append(result, i)
		}
	}

	return result
}

func (row Row) Equals(other Row) bool {
	for i := range other {
		if row[i] != other[i] {
			return false
		}
	}

	return true
}

func (row Row) CountDiff(other Row) int {
	result := 0

	for i := range other {
		if row[i] != other[i] {
			result++
		}
	}

	return result
}

func (matrix Matrix) ForEachChar(function func(c rune)) {
	for _, row := range matrix {
		row.ForEachChar(function)
	}
}

func (matrix Matrix) ForEachCell(function func(x int, y int, c rune)) {
	for i, row := range matrix {
		for j, val := range row {
			function(i, j, val)
		}
	}
}

func (matrix Matrix) Count(ch rune) int {
	result := 0
	matrix.ForEachChar(func(c rune) {
		if c == ch {
			result++
		}
	})

	return result
}

func (matrix Matrix) FindElement(elem rune) Pair {
	for x, row := range matrix {
		for y, cell := range row {
			if cell == elem {
				return Pair{x, y}
			}
		}
	}

	return Pair{-1, -1}
}

func (matrix Matrix) FindElementAll(elem rune) []Pair {
	var result []Pair

	for x, row := range matrix {
		for y, cell := range row {
			if cell == elem {
				result = append(result, Pair{x, y})
			}
		}
	}

	return result
}

func (matrix Matrix) FindElementInRow(rowIndex int, elem rune) int {
	for y, cell := range matrix[rowIndex] {
		if cell == elem {
			return y
		}
	}

	return -1
}

func (matrix Matrix) FindElementInCol(colIndex int, elem rune) int {
	for x, row := range matrix {
		if row[colIndex] == elem {
			return x
		}
	}

	return -1
}

func (matrix Matrix) GetColumn(colIndex int) []rune {
	result := make([]rune, matrix.Len())

	for x, row := range matrix {
		result[x] = row[colIndex]
	}

	return result
}

func (matrix Matrix) Len() int {
	return len(matrix)
}

func (matrix Matrix) Index(i int) Row {
	return matrix[i]
}

func (matrix Matrix) Print() {
	for _, row := range matrix {
		for _, col := range row {
			print(string(col))
		}
		println()
	}
}

func (matrix Matrix) PrintWithHighlightList(cellsToHighlight []Pair) {
	highlightMap := make(map[Pair]bool)

	for _, el := range cellsToHighlight {
		highlightMap[el] = true
	}

	matrix.PrintWithHighlightMap(highlightMap)
}

func (matrix Matrix) PrintWithHighlightMap(cellsToHighlight map[Pair]bool) {
	red := color.New(color.FgRed, color.Bold)
	yellow := color.New(color.FgHiBlue, color.Bold)

	for ir, r := range matrix {
		for ic, e := range r {
			point := Pair{E1: ir, E2: ic}
			if _, ok := cellsToHighlight[point]; ok {
				_, err := red.Print(string(e))
				if err != nil {
					return
				}
			} else {
				_, err := yellow.Print(string(e))
				if err != nil {
					return
				}
			}
		}

		fmt.Println()
	}
}

func (matrix Matrix) Column(index int) Column {
	var result Column

	for _, el := range matrix {
		result = append(result, el[index])
	}

	return result
}

func (matrix Matrix) Transpose() Matrix {
	var result Matrix

	for i := 0; i < len(matrix[0]); i++ {
		result = append(result, Row(matrix.Column(i)))
	}

	return result
}

// Rotate degrees must be divisible by 90
func (matrix Matrix) Rotate(degrees int) Matrix {
	degrees = degrees % 360

	if degrees < 0 {
		degrees += 360
	}

	if degrees == 0 {
		return matrix
	}

	if degrees == 90 {
		return matrix.Transpose().VerticalMirror()
	}

	if degrees == 180 {
		return matrix.Rotate(90).Rotate(90)
	}

	if degrees == 270 {
		return matrix.Rotate(180).Rotate(90)
	}

	return nil
}

func (matrix Matrix) Slice(x1 int, y1 int, x2 int, y2 int) Matrix {
	if x1 >= x2 || y1 >= y2 {
		return nil
	}

	var result Matrix

	for i := x1; i < x2; i++ {
		result = append(result, matrix[i][y1:y2])
	}

	return result
}

func (matrix Matrix) RowSlice(x1 int, x2 int) Matrix {
	return matrix.Slice(x1, 0, x2, len(matrix[0]))
}

func (matrix Matrix) Equals(other Matrix) bool {
	if matrix.Len() != other.Len() {
		return false
	}

	for i := range matrix {
		if !matrix[i].Equals(other[i]) {
			return false
		}
	}

	return true
}

func (matrix Matrix) HorizontalMirror() Matrix {
	var result Matrix

	for i := len(matrix) - 1; i >= 0; i-- {
		result = append(result, matrix[i])
	}

	return result
}

func (matrix Matrix) VerticalMirror() Matrix {
	var result Matrix

	for _, row := range matrix {
		var newRow Row

		for i := len(row) - 1; i >= 0; i-- {
			newRow = append(newRow, row[i])
		}

		result = append(result, newRow)
	}

	return result
}

func (matrix Matrix) CountDiff(other Matrix) int {
	result := 0

	for i := range matrix {
		result += matrix[i].CountDiff(other[i])
	}

	return result
}

func (matrix Matrix) Copy() Matrix {
	var result Matrix

	for _, row := range matrix {
		var newRow Row

		for _, cell := range row {
			newRow = append(newRow, cell)
		}

		result = append(result, newRow)
	}

	return result
}

func (matrix Matrix) Hash() uint64 {
	hash := fnv.New64a()

	for _, row := range matrix {
		for _, char := range row {
			hash.Write([]byte(string(char)))
		}
	}

	return hash.Sum64()
}

func (col Column) Equals(other Column) bool {
	return Row(col).Equals(Row(other))
}

func (col Column) FindElementAll(elem rune) []int {
	return Row(col).FindElementAll(elem)
}
