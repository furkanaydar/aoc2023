package utils

import (
	"regexp"
	"strconv"
	"strings"
)

type String string

func (input String) Index(other string) int {
	return strings.Index(string(input), other)
}

func (input String) IndexAny(other string) int {
	return strings.IndexAny(string(input), other)
}

func (input String) LastIndex(other string) int {
	return strings.LastIndex(string(input), other)
}

func (input String) LastIndexAny(other string) int {
	return strings.LastIndexAny(string(input), other)
}

func (input String) Splitter(sep string) StringArray {
	arr := strings.Split(string(input), sep)
	result := make(StringArray, len(arr))

	for index, elem := range arr {
		result[index] = String(elem)
	}

	return result
}

func (input String) TrimLeft(ch string) String {
	return String(strings.TrimLeft(string(input), ch))
}

func (input String) Contains(other string) bool {
	return strings.Contains(string(input), other)
}

func (input String) ToInt() (int, error) {
	result, err := strconv.Atoi(string(input))

	if err == nil {
		return result, nil
	}

	return 0, err
}

func (input String) ToIntOrDefault(defaultValue int) int {
	result, err := strconv.Atoi(string(input))

	if err == nil {
		return result
	}

	return defaultValue
}

func (input String) IsNumber() bool {
	_, err := input.ToInt()
	return err == nil
}

func (input String) NumbersAsStrings() StringArray {
	spaceSeparatedArr := input.SeparatedBySpace()
	numbersStartingAtIndex := spaceSeparatedArr.NumbersStartingAtIndex()
	return spaceSeparatedArr[numbersStartingAtIndex:]
}

func (input String) NumbersAsInt() []int {
	return input.NumbersAsStrings().ToIntArr()
}

func (input String) GetIndexesOf(find rune) []int {
	var result []int

	for i, ch := range input {
		if ch == find {
			result = append(result, i)
		}
	}

	return result
}

func (input String) SeparatedBySpace() StringArray {

	regex := regexp.MustCompile(`\s+`)
	spacesFormatted := regex.ReplaceAllString(strings.TrimLeft(string(input), " "), " ")
	split := strings.Split(spacesFormatted, " ")
	result := make(StringArray, len(split))
	for index, val := range split {
		result[index] = String(val)
	}

	return result
}

func (input String) SeparatedByComma() StringArray {
	split := strings.Split(string(input), ",")
	result := make(StringArray, len(split))
	for index, val := range split {
		result[index] = String(val)
	}

	return result
}

func (input String) SeparatedBySpaceNSlices(firstNSpace int) StringArray {
	regex := regexp.MustCompile(`\s+`)
	spacesFormatted := regex.ReplaceAllString(string(input), " ")
	split := strings.SplitN(spacesFormatted, " ", firstNSpace)
	result := make(StringArray, len(split))
	for index, val := range split {
		result[index] = String(val)
	}

	return result
}

func (input String) ParenthesesUnwrapped(separator string) StringArray {
	if input.IsFirstElement('(') && input.IsLastElement(')') {
		return input[1 : len(input)-1].Splitter(separator)
	}

	return StringArray{}
}

func (input String) ReplaceAll(i string, o string) String {
	return String(strings.ReplaceAll(string(input), i, o))
}

func (input String) FirstElement() uint8 {
	return input[0]
}

func (input String) IsFirstElement(ch uint8) bool {
	return input.FirstElement() == ch
}

func (input String) LastElement() uint8 {
	return input[len(input)-1]
}

func (input String) IsLastElement(ch uint8) bool {
	return input.LastElement() == ch
}

type StringArray []String

func (input StringArray) Len() int           { return len(input) }
func (input StringArray) Less(i, j int) bool { return input[i] < input[j] }
func (input StringArray) Swap(i, j int)      { input[i], input[j] = input[j], input[i] }

func (input StringArray) NumbersStartingAtIndex() int {
	for index, elem := range input {
		if elem.IsNumber() {
			return index
		}
	}

	return -1
}

func (input StringArray) IndexOf(other String) int {
	for index, elem := range input {
		if elem == other {
			return index
		}
	}

	return -1
}

func (input StringArray) ToIntArr() []int {
	numbersInt := make([]int, len(input))

	for index, val := range input {
		intVal, err := val.ToInt()

		if err == nil {
			numbersInt[index] = intVal
		}

	}

	return numbersInt
}

func (input StringArray) ToMatrix() Matrix {
	var result Matrix

	for _, line := range input {
		result = append(result, Row(line))
	}

	return result
}
