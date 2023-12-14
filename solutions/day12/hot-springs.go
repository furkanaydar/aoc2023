package day12

import (
	"AdventOfCode2023/solutions/utils"
)

var format utils.String
var groups []int
var d [][][]int

func HotSprings1() int {
	input := utils.NewProblem("solutions/day12/input.txt").InputAsLines()
	result := 0

	for _, l := range input {
		separatedBySpace := l.SeparatedBySpace()
		format = separatedBySpace[0]
		groups = separatedBySpace[1].SeparatedByComma().ToIntArr()
		initializeMemo()
		result += waysToFormat(0, 0, 0, false)
	}

	return result
}

func HotSprings2() int {
	input := utils.NewProblem("solutions/day12/input.txt").InputAsLines()
	result := 0

	for _, l := range input {
		separatedBySpace := l.SeparatedBySpace()
		format = ""
		groups = []int{}

		_format := separatedBySpace[0]
		_groups := separatedBySpace[1].SeparatedByComma().ToIntArr()

		for i := 1; i <= 5; i++ {
			format += _format + "?"
		}

		format = format[:len(format)-1]
		for i := 1; i <= 5; i++ {
			groups = append(groups, _groups...)
		}

		initializeMemo()
		result += waysToFormat(0, 0, 0, false)
	}

	return result
}

func waysToFormat(curIndex int, squareCount int, groupId int, finished bool) int {
	if curIndex >= len(format) {
		if finished {
			return 1
		}

		return 0
	}

	if curIndex < len(format) && groupId < len(groups) && d[curIndex][groupId][squareCount] != -1 {
		return d[curIndex][groupId][squareCount]
	}

	if format[curIndex] == '#' {
		if finished {
			setMemo(curIndex, squareCount, groupId, 0)
			return 0
		}

		if squareCount+1 > groups[groupId] {
			setMemo(curIndex, squareCount, groupId, 0)
			return 0
		}

		if squareCount+1 == groups[groupId] && groupId+1 == len(groups) {
			finished = true
		}

		result := waysToFormat(curIndex+1, squareCount+1, groupId, finished)
		setMemo(curIndex, squareCount, groupId, result)
		return result
	}

	if format[curIndex] == '.' {
		if squareCount == 0 {
			result := waysToFormat(curIndex+1, 0, groupId, finished)
			setMemo(curIndex, squareCount, groupId, result)
			return result
		}

		if squareCount != groups[groupId] {
			setMemo(curIndex, squareCount, groupId, 0)
			return 0
		}

		result := waysToFormat(curIndex+1, 0, groupId+1, finished)
		setMemo(curIndex, squareCount, groupId, result)
		return result
	}

	result := 0
	if squareCount == 0 {
		result += waysToFormat(curIndex+1, 0, groupId, finished)
	} else if squareCount == groups[groupId] {
		result += waysToFormat(curIndex+1, 0, groupId+1, finished)
	}

	if !finished {
		if squareCount+1 == groups[groupId] && groupId+1 == len(groups) {
			finished = true
		}

		result += waysToFormat(curIndex+1, squareCount+1, groupId, finished)
	}

	setMemo(curIndex, squareCount, groupId, result)
	return result
}

func initializeMemo() {
	d = make([][][]int, len(format))

	for i := range format {
		d[i] = make([][]int, len(groups))

		for j := range groups {
			d[i][j] = make([]int, 100)
		}
	}

	for i := range format {
		for j := range groups {
			for k := range d[i][j] {
				d[i][j][k] = -1
			}
		}
	}
}

func setMemo(curIndex int, squareCount int, groupId int, val int) {
	if curIndex < len(format) && groupId < len(groups) {
		d[curIndex][groupId][squareCount] = val
	}
}
