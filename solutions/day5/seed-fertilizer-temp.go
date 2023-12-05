package day5

/*
package day5

import (
	"AdventOfCode2023/solutions/utils"
	"sort"
	"strings"
)

func SeedFertilizer1() string {
	problem := utils.Problem{
		InputFileName: "solutions/day5/input.txt",
		Solver: func(input []string) string {
			seeds := readSeeds(input)
			groups := readConversions(input)
			for _, group := range groups {
				sort.Ints(seeds)
				seeds = transform(seeds, group)
			}

			result := seeds[0]
			for _, value := range seeds[1:] {
				if value < result {
					result = value
				}
			}

			return utils.FromIntToString(result)
		},
	}

	return problem.Solve()
}

func SeedFertilizer2() string {
	problem := utils.Problem{
		InputFileName: "solutions/day5/input.txt",
		Solver: func(input []string) string {
			seeds := modifySeeds(readSeeds(input))
			groups := readConversions(input)
			for _, group := range groups {
				sort.Ints(seeds)
				seeds = transform(seeds, group)
			}

			result := seeds[0]
			for _, value := range seeds[1:] {
				if value < result {
					result = value
				}
			}

			return utils.FromIntToString(result)
		},
	}

	return problem.Solve()
}

func readSeeds(input []string) []int {
	return convertNumbersStrToNumbersInt(strings.TrimPrefix(input[0], "seeds: "))
}

func modifySeeds(input []int) []int {
	var result []int

	for i := 0; i < len(input); i++ {
		for j := input[i]; j <= input[i]+input[i+1]-1; j++ {
			result = append(result, j)
		}

		i++
	}

	return result
}

func readConversions(input []string) []SrcDestRangeGroup {
	var result []SrcDestRangeGroup
	var currentGroup SrcDestRangeGroup

	for _, line := range input[2:] {
		if line == "" {
			result = append(result, currentGroup)
			currentGroup = SrcDestRangeGroup{}
		} else if !strings.Contains(line, "map") {
			numbers := convertNumbersStrToNumbersInt(line)
			currentGroup = append(currentGroup,
				SrcDestRange{numbers[0], numbers[1], numbers[1] + numbers[2] - 1})
		}
	}

	return append(result, currentGroup)
}

func transform(input []int, transformer SrcDestRangeGroup) []int {
	sort.Sort(transformer)
	transformerIndex := 0
	result := make([]int, len(input))

	for index, value := range input {

		for value > transformer[transformerIndex].SrcEnd && transformerIndex+1 < transformer.Len() {
			transformerIndex++
		}

		groupStart := transformer[transformerIndex].SrcStart
		groupEnd := transformer[transformerIndex].SrcEnd

		if value < groupStart || value > groupEnd {
			result[index] = value
		} else {
			result[index] = value + transformer[transformerIndex].DestStart - transformer[transformerIndex].SrcStart
		}
	}

	return result
}

func convertNumbersStrToNumbersInt(input string) []int {
	numbersStr := strings.Split(input, " ")
	numbersInt := make([]int, len(numbersStr))

	for index, val := range numbersStr {
		numbersInt[index] = utils.FromStringToInt(val)
	}

	return numbersInt
}

*/
