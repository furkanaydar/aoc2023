package day15

import "AdventOfCode2023/solutions/utils"

type BoxElement struct {
	label       utils.String
	focalLength int
}

type BoxElementList []BoxElement

func (boxes *BoxElementList) RemoveIndex(index int) {
	*boxes = append((*boxes)[:index], (*boxes)[index+1:]...)
}

func (boxes *BoxElementList) Remove(elem BoxElement) {
	for i, el := range *boxes {
		if el.label == elem.label {
			boxes.RemoveIndex(i)
		}
	}
}

func (boxes *BoxElementList) Add(elem BoxElement) {
	*boxes = append(*boxes, elem)
}

func (boxes *BoxElementList) Get(index int) BoxElement {
	return (*boxes)[index]
}

func (boxes *BoxElementList) Replace(index int, elem BoxElement) {
	(*boxes)[index] = elem
}

func (boxes *BoxElementList) AddOrReplace(elem BoxElement) {
	for i, el := range *boxes {
		if el.label == elem.label {
			boxes.Replace(i, elem)
			return
		}
	}

	boxes.Add(elem)
}
