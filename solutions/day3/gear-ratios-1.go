package day3

import (
	"AdventOfCode2023/solutions/utils"
	"strconv"
	"unicode"
)

func GearRatios1() int {
	input := utils.NewProblem("solutions/day3/input.txt").InputAsLines()
	result := 0

	for index, line := range input {
		for it := 0; it < len(line); it++ {
			if unicode.IsDigit(rune(line[it])) {
				var numberStr string
				y1 := it
				y2 := it
				for y2 < len(line) && unicode.IsDigit(rune(line[y2])) {
					numberStr += string(line[y2])
					y2++
				}

				it = y2

				if hasSymbolAround(input, index, y1, y2-1) {
					num, err := strconv.Atoi(numberStr)

					if err != nil {
						return -1
					}

					result += num
				}
			}
		}
	}

	return result
}

func hasSymbolAround(input utils.StringArray, x int, y1 int, y2 int) bool {
	isSymbolIndexSafe := func(r int, c int) bool {
		if r < 0 || r >= len(input) || c < 0 || c >= len(input[r]) {
			return false
		}

		return isSymbol(rune(input[r][c]))
	}

	for _, r := range []int{x - 1, x + 1} {
		for c := y1 - 1; c <= y2+1; c++ {
			if isSymbolIndexSafe(r, c) {
				return true
			}
		}
	}

	return isSymbolIndexSafe(x, y1-1) || isSymbolIndexSafe(x, y2+1)
}

func isSymbol(input rune) bool {
	return !unicode.IsDigit(input) && input != '.'
}
