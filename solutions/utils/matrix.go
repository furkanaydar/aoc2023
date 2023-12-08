package utils

type Cell struct {
	X   int
	Y   int
	val rune
}

type Row []rune

func (row Row) ForEachChar(function func(c rune)) {
	for _, ch := range row {
		function(row[ch])
	}
}

func (row Row) ForEachCell(function func(i int, c rune)) {
	for index, ch := range row {
		function(index, ch)
	}
}

type Matrix []Row

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
