package day2

import (
	"AdventOfCode2023/solutions/utils"
	"fmt"
	"strconv"
	"strings"
)

type ColorCount struct {
	Blue  int
	Green int
	Red   int
}

func parseGame(line string) []ColorCount {
	startIndex := strings.Index(line, ": ") + 2
	line = line[startIndex:]
	sliced := strings.Split(line, "; ")
	result := make([]ColorCount, len(sliced))

	for tryIndex, colors := range sliced {
		pickedColors := strings.Split(colors, ",")

		for _, countAndColor := range pickedColors {
			countAndColor = strings.TrimLeft(countAndColor, " ")

			countAndColorArr := strings.Split(countAndColor, " ")
			pickedCount, err := strconv.Atoi(countAndColorArr[0])
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

func parseGames() [][]ColorCount {
	problem := utils.Problem{InputFileName: "solutions/day2/input.txt"}
	input := problem.ReadInputToLines()

	allGames := make([][]ColorCount, len(input))

	for index, game := range input {
		allGames[index] = parseGame(game)
	}

	return allGames
}

func CubeConundrum1() int {
	result := 0

	var ColorMax = map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}

	for index, game := range parseGames() {
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
	result := 0

	for _, game := range parseGames() {
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