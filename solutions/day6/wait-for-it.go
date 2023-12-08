package day6

import (
	"AdventOfCode2023/solutions/utils"
)

func WaitForIt1() int {
	input := utils.NewProblem("solutions/day6/input.txt").InputAsLines()
	times := input[0].NumbersAsInt()
	distances := input[1].NumbersAsInt()
	result := 1

	for index, time := range times {
		count := 0
		for i := 1; i <= time-1; i++ {
			if (time-i)*i > distances[index] {
				count++
			}
		}

		result *= count
	}

	return result
}

func WaitForIt2() int {
	input := utils.NewProblem("solutions/day6/input.txt").InputAsLines()
	times := []int{readLineNumber(input[0])}
	distances := []int{readLineNumber(input[1])}
	result := 1

	for index, time := range times {
		count := 0
		for i := 1; i <= time-1; i++ {
			if (time-i)*i > distances[index] {
				count++
			}
		}

		result *= count
	}

	return result
}

func readLineNumber(input utils.String) int {
	elems := input.SeparatedBySpace()
	result := ""

	for _, toAdd := range elems[elems.NumbersStartingAtIndex():] {
		result += string(toAdd)
	}

	return utils.String(result).ToIntOrDefault(-1)
}
