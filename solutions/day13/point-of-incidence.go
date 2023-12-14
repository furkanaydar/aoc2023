package day13

import (
	"AdventOfCode2023/solutions/utils"
	"fmt"
)

func PointOfIncidence1() int {
	return Solve(func(top utils.Matrix, bottom utils.Matrix) bool {
		return top.Equals(bottom)
	})
}

func PointOfIncidence2() int {
	return Solve(func(top utils.Matrix, bottom utils.Matrix) bool {
		return top.CountDiff(bottom) == 1
	})
}

func Solve(validator func(utils.Matrix, utils.Matrix) bool) int {
	result := 0
	for _, matrix := range readAllMatrices() {
		horizontalScore := pointForMatrix(matrix, validator) * 100
		verticalScore := pointForMatrix(matrix.Transpose(), validator)
		result += horizontalScore + verticalScore
		fmt.Println()
	}

	return result
}

func readAllMatrices() []utils.Matrix {
	input := utils.NewProblem("solutions/day13/input.txt").InputAsLines()
	var matrices []utils.Matrix
	var currentCase utils.StringArray

	for _, l := range input {
		if len(l) == 0 {
			matrices = append(matrices, currentCase.ToMatrix())
			currentCase = utils.StringArray{}
		} else {
			currentCase = append(currentCase, l)
		}
	}

	matrices = append(matrices, currentCase.ToMatrix())
	return matrices
}

func pointForMatrix(input utils.Matrix, validator func(utils.Matrix, utils.Matrix) bool) int {
	result := 0
	inputLen := len(input)

	for i := 0; i < inputLen-1; i++ {
		iterate := min(i+1, inputLen-i-1)
		top := input.RowSlice(i-iterate+1, i+1)
		bottom := input.RowSlice(i+1, i+iterate+1)

		if validator(top.HorizontalMirror(), bottom) {
			result += i + 1
		}
	}

	return result
}
