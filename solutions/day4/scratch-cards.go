package day4

import (
	"AdventOfCode2023/solutions/utils"
	"math"
)

func ScratchCards1() int {
	input := utils.NewProblem("solutions/day4/input.txt").InputAsLines()
	result := 0

	resultCalculator := func(_ int, matches int) {
		if matches > 0 {
			result += int(math.Pow(2, float64(matches-1)))
		}
	}

	applyToMatchingNumbersForEveryCard(input, resultCalculator)
	return result
}

func ScratchCards2() int {
	input := utils.NewProblem("solutions/day4/input.txt").InputAsLines()
	cardCounts := make(map[int]int)

	for i := 1; i <= len(input); i++ {
		cardCounts[i] = 1
	}

	resultCalculator := func(index int, matches int) {
		for i := index + 2; i <= index+1+matches; i++ {
			cur := cardCounts[i]
			cardCounts[i] = cur + cardCounts[index+1]
		}
	}

	applyToMatchingNumbersForEveryCard(input, resultCalculator)

	result := 0
	for _, val := range cardCounts {
		result += val
	}

	return result
}

func applyToMatchingNumbersForEveryCard(inputLines utils.StringArray, calculate func(int, int)) {
	for index, line := range inputLines {
		selectedNumbers, winningNumbers := getSelectedAndWinningNumbers(line)
		winningNumberSet := make(map[int]bool)

		for _, winningNumber := range winningNumbers {
			winningNumberSet[winningNumber] = true
		}

		matches := 0

		for _, selectedNumber := range selectedNumbers {
			if winningNumberSet[selectedNumber] {
				matches++
			}
		}

		calculate(index, matches)
	}
}

func getSelectedAndWinningNumbers(line utils.String) ([]int, []int) {
	numbersStartingAt := line.Index(": ") + 2
	numbers := line[numbersStartingAt:].Splitter("|")
	selectedNumbersFiltered := numbers[0].SeparatedBySpace().ToIntArr()
	winningNumbersFiltered := numbers[1].SeparatedBySpace().ToIntArr()
	return selectedNumbersFiltered, winningNumbersFiltered
}
