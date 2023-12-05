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
			seeds := convertNumbersStrToNumbersInt(strings.TrimPrefix(input[0], "seeds: "))
			seedsRange := make([]Range, len(seeds))

			for index, seed := range seeds {
				seedsRange[index] = Range{Start: seed, End: seed}
			}

			return Solve(input, seedsRange)
		},
	}

	return problem.Solve()
}

func SeedFertilizer2() string {
	problem := utils.Problem{
		InputFileName: "solutions/day5/input.txt",
		Solver: func(input []string) string {
			seeds := convertNumbersStrToNumbersInt(strings.TrimPrefix(input[0], "seeds: "))
			seedsRange := make([]Range, len(seeds)/2)

			for i := 0; i < len(seeds); i += 2 {
				seedsRange[i/2] = Range{Start: seeds[i], End: seeds[i] + seeds[i+1]}
			}

			return Solve(input, seedsRange)
		},
	}

	return problem.Solve()
}

func Solve(input []string, seedsRange RangeGroup) string {
	groups := readConversions(input)

	for _, group := range groups {
		sort.Sort(group)
		sort.Sort(seedsRange)

		for _, transformer := range group {
			var tmpRanges []Range

			for _, seedRange := range seedsRange {
				splitResult := transformer.SrcRange.splitRange(seedRange)
				tmpRanges = append(tmpRanges, splitResult...)
			}

			seedsRange = tmpRanges
		}

		hash := make([]bool, seedsRange.Len()*2)
		for _, transformer := range group {
			var tmpRanges []Range

			for index, seedRange := range seedsRange {
				compareResult := transformer.SrcRange.compareRange(seedRange)

				if !hash[index] && compareResult == InRange {
					tmpRanges = append(tmpRanges, applyTransformerToSeedRange(transformer, seedRange))
					hash[index] = true
				} else {
					tmpRanges = append(tmpRanges, seedRange)
				}
			}

			seedsRange = tmpRanges
		}
	}

	result := seedsRange[0].Start
	for _, value := range seedsRange[1:] {
		if value.Start < result {
			result = value.Start
		}
	}

	return utils.FromIntToString(result)
}

func applyTransformerToSeedRange(transformerRange SrcDestRange, seedRange Range) Range {
	return Range{
		Start: transformerRange.DestStart + seedRange.Start - transformerRange.SrcRange.Start,
		End:   transformerRange.DestStart + seedRange.End - transformerRange.SrcRange.Start,
	}
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
				SrcDestRange{numbers[0], Range{Start: numbers[1], End: numbers[1] + numbers[2] - 1}},
			)
		}
	}

	return append(result, currentGroup)
}

func convertNumbersStrToNumbersInt(input string) []int {
	numbersStr := strings.Split(input, " ")
	numbersInt := make([]int, len(numbersStr))

	for index, val := range numbersStr {
		numbersInt[index] = utils.FromStringToInt(val)
	}

	return numbersInt
}
