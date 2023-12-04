package day4

import (
	"AdventOfCode2023/solutions/utils"
	"math"
	"strconv"
	"strings"
)

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

func getSelectedAndWinningNumbers(line string) ([]int, []int) {
	numbersStartingAt := strings.Index(line, ": ") + 2
	numbers := strings.Split(line[numbersStartingAt:], "|")
	selectedNumbersFiltered := filterNumbers(strings.Split(numbers[0], " "))
	winningNumbersFiltered := filterNumbers(strings.Split(numbers[1], " "))
	return selectedNumbersFiltered, winningNumbersFiltered
}

func applyToMatchingNumbersForEveryCard(inputLines []string, calculate func(int, int)) {
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

func ScratchCards1() int {
	problem := utils.Problem{InputFileName: "solutions/day4/input.txt"}
	inputLines := problem.ReadInputToLines()
	result := 0

	resultCalculator := func(_ int, matches int) {
		if matches > 0 {
			result += int(math.Pow(2, float64(matches-1)))
		}
	}

	applyToMatchingNumbersForEveryCard(inputLines, resultCalculator)
	return result
}

func ScratchCards2() int {
	problem := utils.Problem{InputFileName: "solutions/day4/input.txt"}
	inputLines := problem.ReadInputToLines()
	cardCounts := make(map[int]int)

	for i := 1; i <= len(inputLines); i++ {
		cardCounts[i] = 1
	}

	resultCalculator := func(index int, matches int) {
		for i := index + 2; i <= index+1+matches; i++ {
			cur := cardCounts[i]
			cardCounts[i] = cur + cardCounts[index+1]
		}
	}

	applyToMatchingNumbersForEveryCard(inputLines, resultCalculator)

	result := 0
	for _, val := range cardCounts {
		result += val
	}

	return result
}