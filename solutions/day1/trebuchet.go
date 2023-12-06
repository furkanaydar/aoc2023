package day1

import (
	"AdventOfCode2023/solutions/utils"
)

func TrebuchetPart1() string {
	problem := utils.Problem{
		InputFileName: "solutions/day1/input.txt",
		Solver: func(input utils.AocStringArray) utils.Any {
			result := 0

			for _, line := range input {
				first, _ := captureFirstDigitIndexFromDigit(line)
				last, _ := captureLastDigitIndexFromDigit(line)
				result += int(first)*10 + int(last)
			}

			return result
		},
	}

	return problem.Solve()
}

func TrebuchetPart2() string {
	problem := utils.Problem{
		InputFileName: "solutions/day1/input.txt",
		Solver: func(input utils.AocStringArray) utils.Any {
			result := 0

			for _, line := range input {
				first := captureFirstDigit(line)
				last := captureLastDigit(line)
				result += int(first)*10 + last
			}

			return result
		},
	}

	return problem.Solve()
}

const chars = "0123456789"

var digitTexts = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func captureFirstDigitIndexFromText(line utils.AocString) (int, int) {
	minIndex := 100000
	digit := -1

	for textIndex, text := range digitTexts {
		index := line.Index(text)
		if index != -1 && index < minIndex {
			minIndex = index
			digit = textIndex
		}
	}

	return digit, minIndex
}

func captureLastDigitIndexFromText(line utils.AocString) (int, int) {
	maxIndex := -1
	digit := -1

	for textIndex, text := range digitTexts {
		index := line.LastIndex(text)
		if index != -1 && index > maxIndex {
			maxIndex = index
			digit = textIndex
		}
	}

	return digit, maxIndex
}

func captureFirstDigitIndexFromDigit(line utils.AocString) (uint8, int) {
	firstDigitCharacterIndex := line.IndexAny(chars)
	return line[firstDigitCharacterIndex] - '0', firstDigitCharacterIndex
}

func captureLastDigitIndexFromDigit(line utils.AocString) (uint8, int) {
	lastDigitCharacterIndex := line.LastIndexAny(chars)
	return line[lastDigitCharacterIndex] - '0', lastDigitCharacterIndex
}

func captureFirstDigit(line utils.AocString) int {
	asDigitValue, asDigitIndex := captureFirstDigitIndexFromDigit(line)
	asTextValue, asTextIndex := captureFirstDigitIndexFromText(line)
	if asDigitIndex < asTextIndex {
		return int(asDigitValue)
	}

	return asTextValue
}

func captureLastDigit(line utils.AocString) int {
	asDigitValue, asDigitIndex := captureLastDigitIndexFromDigit(line)
	asTextValue, asTextIndex := captureLastDigitIndexFromText(line)

	if asDigitIndex > asTextIndex {
		return int(asDigitValue)
	}

	return asTextValue
}
