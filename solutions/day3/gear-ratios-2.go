package day3

import (
	"AdventOfCode2023/solutions/utils"
	"strconv"
	"unicode"
)

func GearRatios2() string {
	problem := utils.Problem{
		InputFileName: "solutions/day3/input.txt",
		Solver: func(input utils.AocStringArray) utils.Any {
			stars := make(map[utils.Cell][]int)

			for x, line := range input {
				for y := 0; y < len(line); y++ {
					if unicode.IsDigit(rune(line[y])) {
						it := y
						var numberStr string

						for it < len(line) && unicode.IsDigit(rune(line[it])) {
							numberStr += string(line[it])
							it++
						}

						actualNumber, err := strconv.Atoi(numberStr)

						if err == nil {
							solveForNumber(input, x, y, it-1, stars, actualNumber)
						}

						y = it
					}
				}
			}

			result := 0

			for key := range stars {
				if len(stars[key]) > 1 {
					ratio := 1
					for _, val := range stars[key] {
						ratio *= val
					}
					result += ratio
				}
			}

			return result
		},
	}

	return problem.Solve()
}

func solveForNumber(input utils.AocStringArray, x int, y1 int, y2 int, stars map[utils.Cell][]int, actualNumber int) {
	isStar := func(r int, c int) bool {
		return input[r][c] == '*'
	}

	indexSafeValidator := func(r int, c int) bool {
		return r >= 0 && r < len(input) && c >= 0 && c < len(input[r]) && isStar(r, c)
	}

	appendToMap := func(r int, c int) {
		curCell := utils.Cell{
			X: r, Y: c,
		}

		stars[curCell] = append(stars[curCell], actualNumber)
	}

	for _, r := range []int{x - 1, x + 1} {
		for c := y1 - 1; c <= y2+1; c++ {
			if indexSafeValidator(r, c) {
				appendToMap(r, c)
			}
		}
	}

	if indexSafeValidator(x, y1-1) {
		appendToMap(x, y1-1)
	}

	if indexSafeValidator(x, y2+1) {
		appendToMap(x, y2+1)
	}
}
