package day11

import (
	"AdventOfCode2023/solutions/utils"
	"math"
)

func CosmicExpansion1() int {
	return solve(1)
}

func CosmicExpansion2() int {
	return solve(999999)
}

func solve(expansionFactor int) int {
	input := utils.NewProblem("solutions/day11/input.txt").InputAsMatrix()

	var noGalaxyRows []int
	var noGalaxyCols []int

	for i, r := range input {
		if r.FindElement('#') == -1 {
			noGalaxyRows = append(noGalaxyRows, i)
		}
	}

	for i := range input[0] {
		if input.FindElementInCol(i, '#') == -1 {
			noGalaxyCols = append(noGalaxyCols, i)
		}
	}

	galaxies := input.FindElementAll('#')
	result := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1x := galaxies[i].E1
			g1y := galaxies[i].E2
			g2x := galaxies[j].E1
			g2y := galaxies[j].E2
			dist := int(math.Abs(float64(g1x-g2x)) + math.Abs(float64(g1y-g2y)))

			for k := min(g1x, g2x); k <= max(g1x, g2x); k++ {
				if arrContainsElem(noGalaxyRows, k) {
					dist += expansionFactor
				}
			}

			for k := min(g1y, g2y); k <= max(g1y, g2y); k++ {
				if arrContainsElem(noGalaxyCols, k) {
					dist += expansionFactor
				}
			}

			result += dist
		}
	}
	return result
}

func arrContainsElem(arr []int, elem int) bool {
	for _, el := range arr {
		if el == elem {
			return true
		}
	}

	return false
}
