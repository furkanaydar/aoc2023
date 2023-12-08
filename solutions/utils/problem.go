package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	InputFileName string
	InputLines    StringArray
	Solver        func(input Input) Any
	Elapsed       time.Duration
}

func NewProblem(inputFileName string) *Problem {
	return &Problem{
		InputLines: readInputToLines(inputFileName),
	}
}

type Input StringArray

func (problem Problem) InputAsLines() StringArray {
	return problem.InputLines
}

func (problem Problem) InputAsMatrix() Matrix {
	var result Matrix

	for index, line := range problem.InputLines {
		result[index] = Row(line)
	}

	return result
}

func readInputToLines(inputFileName string) StringArray {
	file, err := os.Open(inputFileName)

	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines StringArray
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lines = append(lines, String(fileScanner.Text()))
	}

	err = fileScanner.Err()

	if err != nil {
		fmt.Println(err)
	}

	return lines
}
