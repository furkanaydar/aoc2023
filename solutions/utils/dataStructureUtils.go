package utils

import (
	"strconv"
)

type Cell struct {
	X int
	Y int
}

func FromStringToInt(input string) int {
	result, err := strconv.Atoi(input)

	if err != nil {
		return result
	}

	return -1
}

func FromIntToString(input int) string {
	return strconv.Itoa(input)
}
