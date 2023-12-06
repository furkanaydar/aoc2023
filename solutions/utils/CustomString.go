package utils

import (
	"regexp"
	"strconv"
	"strings"
)

type AocString string

func (input AocString) Index(other string) int {
	return strings.Index(string(input), other)
}

func (input AocString) IndexAny(other string) int {
	return strings.IndexAny(string(input), other)
}

func (input AocString) LastIndex(other string) int {
	return strings.LastIndex(string(input), other)
}

func (input AocString) LastIndexAny(other string) int {
	return strings.LastIndexAny(string(input), other)
}

func (input AocString) Splitter(sep string) []string {
	return strings.Split(string(input), sep)
}

func (input AocString) Contains(other string) bool {
	return strings.Contains(string(input), other)
}

func (input AocString) ToInt() (int, error) {
	result, err := strconv.Atoi(string(input))

	if err == nil {
		return result, nil
	}

	return 0, err
}

func (input AocString) ToIntOrDefault(defaultValue int) int {
	result, err := strconv.Atoi(string(input))

	if err == nil {
		return result
	}

	return defaultValue
}

func (input AocString) IsNumber() bool {
	_, err := input.ToInt()
	return err == nil
}

func (input AocString) NumbersAsStrings() AocStringArray {
	spaceSeparatedArr := input.SeparatedBySpace()
	numbersStartingAtIndex := spaceSeparatedArr.NumbersStartingAtIndex()
	return spaceSeparatedArr[numbersStartingAtIndex:]
}

func (input AocString) NumbersAsInt() []int {
	return input.NumbersAsStrings().ToIntArr()
}

func (input AocString) SeparatedBySpace() AocStringArray {
	regex := regexp.MustCompile(`\s+`)
	spacesFormatted := regex.ReplaceAllString(string(input), " ")
	split := strings.Split(spacesFormatted, " ")
	result := make(AocStringArray, len(split))
	for index, val := range split {
		result[index] = AocString(val)
	}

	return result
}

type AocStringArray []AocString

func (input AocStringArray) Len() int           { return len(input) }
func (input AocStringArray) Less(i, j int) bool { return input[i] < input[j] }
func (input AocStringArray) Swap(i, j int)      { input[i], input[j] = input[j], input[i] }

func (input AocStringArray) NumbersStartingAtIndex() int {
	for index, elem := range input {
		if elem.IsNumber() {
			return index
		}
	}

	return -1
}

func (input AocStringArray) ToIntArr() []int {
	numbersInt := make([]int, len(input))

	for index, val := range input {
		intVal, err := val.ToInt()

		if err != nil {
			return nil
		}

		numbersInt[index] = intVal
	}

	return numbersInt
}
