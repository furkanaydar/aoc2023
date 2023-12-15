package day15

import (
	"AdventOfCode2023/solutions/utils"
)

func hash(input utils.String) int {
	res := 0
	for _, v := range input {
		res += int(v)
		res *= 17
		res %= 256
	}

	return res
}

func LensLibrary1() int {
	input := utils.NewProblem("solutions/day15/input.txt").InputAsLines()[0].SeparatedByComma()
	result := 0

	for _, str := range input {
		result += hash(str)
	}

	return result
}

func LensLibrary2() int {
	input := utils.NewProblem("solutions/day15/input.txt").InputAsLines()[0].SeparatedByComma()
	boxes := make(map[int]*BoxElementList)

	for _, elem := range input {

		if elem.Contains("=") {
			parts := elem.Splitter("=")
			label := parts[0]
			boxId := hash(label)
			length, _ := parts[1].ToInt()
			newElem := BoxElement{label: label, focalLength: length}

			if _, ok := boxes[boxId]; !ok {
				boxes[boxId] = &BoxElementList{newElem}
			} else {
				list := boxes[boxId]
				list.AddOrReplace(newElem)
			}
		} else if elem.Contains("-") {
			minusIndex := elem.Index("-")
			label := elem[:minusIndex]
			boxId := hash(label)
			list := boxes[boxId]

			if list != nil {
				list.Remove(BoxElement{label: label})
			}
		}
	}

	result := 0
	for i, k := range boxes {
		for j, v := range *k {
			result += (i + 1) * (j + 1) * v.focalLength
		}
	}

	return result
}
