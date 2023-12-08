package day8

import (
	"AdventOfCode2023/solutions/utils"
)

func HauntedWasteland1() int {
	input := utils.NewProblem("solutions/day8/input.txt").InputAsLines()
	return solveForStartingPoint(input[0], getPaths(input), "AAA", utils.StringArray{"ZZZ"})
}

func HauntedWasteland2() int {
	input := utils.NewProblem("solutions/day8/input.txt").InputAsLines()
	paths := getPaths(input)
	directions := input[0]
	var startingNodes utils.StringArray
	var endNodes utils.StringArray
	var allResult []int

	for node := range paths {
		if node.IsLastElement('A') {
			startingNodes = append(startingNodes, node)
		}

		if node.IsLastElement('Z') {
			endNodes = append(endNodes, node)
		}
	}

	for _, node := range startingNodes {
		allResult = append(allResult, solveForStartingPoint(directions, paths, node, endNodes))
	}

	var lcm = allResult[0]
	for _, result := range allResult[1:] {
		lcm = utils.LCM(lcm, result)
	}

	return lcm
}

func getPaths(lines utils.StringArray) map[utils.String]utils.StringArray {
	result := make(map[utils.String]utils.StringArray)
	startPoint := utils.String("")

	for _, lookup := range lines[2:] {
		spaceSeparated := lookup.SeparatedBySpaceNSlices(3)
		if startPoint == "" {
			startPoint = spaceSeparated[0]
		}

		result[spaceSeparated[0]] = spaceSeparated[2].ParenthesesUnwrapped(", ")
	}

	return result
}

func solveForStartingPoint(directions utils.String,
	paths map[utils.String]utils.StringArray,
	startingNode utils.String,
	endNodes utils.StringArray) int {
	directionIndex := 0
	curIndex := startingNode

	for endNodes.IndexOf(curIndex) == -1 {
		newIndex := utils.String("")

		if directions[directionIndex%(len(directions))] == 'L' {
			newIndex = paths[curIndex][0]
		} else {
			newIndex = paths[curIndex][1]
		}

		curIndex = newIndex
		directionIndex++
	}

	return directionIndex
}
