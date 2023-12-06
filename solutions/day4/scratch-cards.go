package day4

import (
	"AdventOfCode2023/solutions/utils"
	"math"
	"strconv"
	"strings"
)

func ScratchCards1() string {
	problem := utils.Problem{
		InputFileName: "solutions/day4/input.txt",
		Solver: func(input utils.AocStringArray) utils.Any {
			result := 0

			resultCalculator := func(_ int, matches int) {
				if matches > 0 {
					result += int(math.Pow(2, float64(matches-1)))
				}
			}

			applyToMatchingNumbersForEveryCard(input, resultCalculator)
			return result
		},
	}

	return problem.Solve()
}

func ScratchCards2() string {
	problem := utils.Problem{
		InputFileName: "solutions/day4/input.txt",
		Solver: func(input utils.AocStringArray) utils.Any {
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
		},
	}

	return problem.Solve()
}

func applyToMatchingNumbersForEveryCard(inputLines utils.AocStringArray, calculate func(int, int)) {
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

func getSelectedAndWinningNumbers(line utils.AocString) ([]int, []int) {
	numbersStartingAt := line.Index(": ") + 2
	numbers := line[numbersStartingAt:].Splitter("|")
	selectedNumbersFiltered := filterNumbers(strings.Split(numbers[0], " "))
	winningNumbersFiltered := filterNumbers(strings.Split(numbers[1], " "))
	return selectedNumbersFiltered, winningNumbersFiltered
}

func filterNumbers(input []string) []int {
	var result []int
	for _, numbersStr := range input {
		if strings.TrimSpace(numbersStr) != "" {
			numberAsInt, err := strconv.Atoi(numbersStr)
			if err == nil {
				result = append(result, numberAsInt)
			}
		}
	}

	return result
}
