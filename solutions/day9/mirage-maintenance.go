package day9

import (
	"AdventOfCode2023/solutions/utils"
)

func MirageMaintenance1() int {
	input := utils.NewProblem("solutions/day9/input.txt").InputAsLines()
	result := 0

	for _, elem := range input {
		diffMatrix := getDiffMatrix(elem.SeparatedBySpace().ToIntArr())
		for _, r := range diffMatrix {
			result += r[len(r)-1]
		}
	}

	return result
}

func MirageMaintenance2() int {
	input := utils.NewProblem("solutions/day9/input.txt").InputAsLines()
	result := 0

	for _, elem := range input {
		diffMatrix := getDiffMatrix(elem.SeparatedBySpace().ToIntArr())
		curAddition := diffMatrix[len(diffMatrix)-1][0]

		for index := len(diffMatrix) - 2; index >= -0; index-- {
			curAddition = diffMatrix[index][0] - curAddition
		}

		result += curAddition
	}

	return result
}

func getDiffRow(row []int) []int {
	result := make([]int, len(row)-1)

	for index := 0; index < len(row)-1; index++ {
		result[index] = row[index+1] - row[index]
	}

	return result
}

func oneDistinctElement(row []int) bool {
	elems := make(map[int]bool)

	for _, elem := range row {
		elems[elem] = true
	}

	return len(elems) == 1
}

func getDiffMatrix(row []int) [][]int {
	var result [][]int

	current := row
	for !oneDistinctElement(current) {
		result = append(result, current)
		current = getDiffRow(current)
	}

	result = append(result, current)
	return result
}
