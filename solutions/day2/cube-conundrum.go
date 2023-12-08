package day2

import (
	"AdventOfCode2023/solutions/utils"
	"fmt"
)

type ColorCount struct {
	Blue  int
	Green int
	Red   int
}

func CubeConundrum1() int {
	input := utils.NewProblem("solutions/day2/input.txt").InputAsLines()
	result := 0

	var ColorMax = map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}

	for index, game := range parseGames(input) {
		possibleGame := true
		for _, item := range game {
			if item.Red > ColorMax["red"] || item.Blue > ColorMax["blue"] || item.Green > ColorMax["green"] {
				possibleGame = false
			}
		}

		if possibleGame {
			result += index + 1
		}
	}

	return result
}

func CubeConundrum2() int {
	input := utils.NewProblem("solutions/day2/input.txt").InputAsLines()
	result := 0

	for _, game := range parseGames(input) {
		maxCounts := ColorCount{0, 0, 0}

		for _, item := range game {
			if item.Red > maxCounts.Red {
				maxCounts.Red = item.Red
			}

			if item.Blue > maxCounts.Blue {
				maxCounts.Blue = item.Blue
			}

			if item.Green > maxCounts.Green {
				maxCounts.Green = item.Green
			}
		}

		result += maxCounts.Red * maxCounts.Blue * maxCounts.Green
	}

	return result
}

func parseGames(input utils.StringArray) [][]ColorCount {
	allGames := make([][]ColorCount, len(input))

	for index, game := range input {
		allGames[index] = parseGame(game)
	}

	return allGames
}

func parseGame(line utils.String) []ColorCount {
	startIndex := line.Index(": ") + 2
	line = line[startIndex:]
	sliced := line.Splitter("; ")
	result := make([]ColorCount, len(sliced))

	for tryIndex, colors := range sliced {
		pickedColors := colors.Splitter(",")

		for _, countAndColor := range pickedColors {
			countAndColor = countAndColor.TrimLeft(" ")
			countAndColorArr := countAndColor.Splitter(" ")
			pickedCount, err := countAndColorArr[0].ToInt()
			pickedColor := countAndColorArr[1]

			if err != nil {
				fmt.Println(err)
				return nil
			}

			if pickedColor == "red" {
				result[tryIndex].Red = pickedCount
			} else if pickedColor == "blue" {
				result[tryIndex].Blue = pickedCount
			} else {
				result[tryIndex].Green = pickedCount
			}
		}
	}

	return result
}
