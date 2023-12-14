package day14

import (
	"AdventOfCode2023/solutions/utils"
	"sort"
)

func TiltNorth(input utils.Matrix) utils.Matrix {
	for i := 0; i < len(input[0]); i++ {
		col := input.Column(i)
		rocks := col.FindElementAll('O')
		cubes := col.FindElementAll('#')

		sort.Ints(rocks)
		sort.Ints(cubes)

		nextPlacement := 0
		for _, rock := range rocks {
			placement := nextPlacement
			for _, cube := range cubes {
				if cube < rock {
					placement = max(placement, cube+1)
				} else {
					break
				}
			}

			if input[placement][i] != 'O' {
				input[placement][i] = 'O'
				input[rock][i] = '.'
			}

			nextPlacement = placement + 1
		}
	}

	return input
}

func ParabolicReflectorDish1() int {
	input := TiltNorth(utils.NewProblem("solutions/day14/input.txt").InputAsMatrix())
	result := 0

	for i := 0; i < len(input[0]); i++ {
		col := input.Column(i)
		rocks := col.FindElementAll('O')
		for _, rock := range rocks {
			result += input.Len() - rock
		}
	}

	return result
}

func score(input utils.Matrix) int {
	result := 0
	for j := 0; j < len(input[0]); j++ {
		col := input.Column(j)
		rocks := col.FindElementAll('O')
		for _, rock := range rocks {
			result += input.Len() - rock
		}
	}

	return result
}

func fullTiltCycle(input utils.Matrix) utils.Matrix {
	tiltedNorth := TiltNorth(input.Copy())
	tiltedWest := TiltNorth(tiltedNorth.Rotate(90)).Rotate(270)
	tiltedSouth := TiltNorth(tiltedWest.Rotate(180)).Rotate(180)
	tiltedEast := TiltNorth(tiltedSouth.Rotate(270)).Rotate(90)

	return tiltedEast
}

func ParabolicReflectorDish2() int {
	input := utils.NewProblem("solutions/day14/input.txt").InputAsMatrix()
	m := make(map[uint64]int)
	results := make(map[int]int)
	firstCollision := 0
	rotation := 0

	for i := 1; i <= 3000; i++ {
		tilted := fullTiltCycle(input)
		hash := tilted.Hash()
		input = tilted

		if m[hash] != 0 {
			rotation = i - m[hash]
			results[i] = score(input)
			firstCollision = i
		} else {
			m[hash] = i
		}

		results[i] = score(input)

		if firstCollision != 0 && (1000000000-i)%rotation == 0 {
			return results[i]
		}
	}

	return 0
}
