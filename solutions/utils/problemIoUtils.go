package utils

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct {
	InputFileName string
	Solver        func([]string) string
}

func (problem *Problem) Solve() string {
	return problem.Solver(problem.readInputToLines())
}

func (problem *Problem) readInputToLines() []string {
	file, err := os.Open(problem.InputFileName)

	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines []string
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	err = fileScanner.Err()

	if err != nil {
		fmt.Println(err)
	}

	return lines
}
