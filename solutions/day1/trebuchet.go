package day1

import (
	"AdventOfCode2023/solutions/utils"
	"fmt"
	"strings"
)

const chars = "0123456789"

var digitTexts = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func captureFirstDigitIndexFromText(line string) (int, int) {
	minIndex := 100000
	digit := -1

	for textIndex, text := range digitTexts {
		index := strings.Index(line, text)
		if index != -1 && index < minIndex {
			minIndex = index
			digit = textIndex
		}
	}

	return digit, minIndex
}

func captureLastDigitIndexFromText(line string) (int, int) {
	maxIndex := -1
	digit := -1

	for textIndex, text := range digitTexts {
		index := strings.LastIndex(line, text)
		if index != -1 && index > maxIndex {
			maxIndex = index
			digit = textIndex
		}
	}

	return digit, maxIndex
}

func captureFirstDigitIndexFromDigit(line string) (uint8, int) {
	firstDigitCharacterIndex := strings.IndexAny(line, chars)
	return line[firstDigitCharacterIndex] - '0', firstDigitCharacterIndex
}

func captureLastDigitIndexFromDigit(line string) (uint8, int) {
	lastDigitCharacterIndex := strings.LastIndexAny(line, chars)
	return line[lastDigitCharacterIndex] - '0', lastDigitCharacterIndex
}

func captureFirstDigit(line string) int {
	asDigitValue, asDigitIndex := captureFirstDigitIndexFromDigit(line)
	asTextValue, asTextIndex := captureFirstDigitIndexFromText(line)
	if asDigitIndex < asTextIndex {
		return int(asDigitValue)
	}

	return asTextValue
}

func captureLastDigit(line string) int {
	asDigitValue, asDigitIndex := captureLastDigitIndexFromDigit(line)
	asTextValue, asTextIndex := captureLastDigitIndexFromText(line)

	if asDigitIndex > asTextIndex {
		return int(asDigitValue)
	}

	return asTextValue
}

func TrebuchetPart1() int {
	problem := utils.Problem{InputFileName: "solutions/day1/input.txt"}
	input := problem.ReadInputToLines()

	if input == nil {
		return -1
	}

	result := 0

	for _, line := range input {
		first, _ := captureFirstDigitIndexFromDigit(line)
		last, _ := captureLastDigitIndexFromDigit(line)
		result += int(first)*10 + int(last)
	}

	return result
}

func TrebuchetPart2() int {
	problem := utils.Problem{InputFileName: "solutions/day1/input.txt"}
	input := problem.ReadInputToLines()

	if input == nil {
		fmt.Println("Could not read input file.")
		return -1
	}

	result := 0

	for _, line := range input {
		first := captureFirstDigit(line)
		last := captureLastDigit(line)
		result += int(first)*10 + last
	}

	return result
}
