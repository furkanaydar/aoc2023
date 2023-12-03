package utils

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct {
	InputFileName string
}

func (problem *Problem) ReadInputToLines() []string {
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
