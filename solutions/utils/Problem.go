package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Problem struct {
	InputFileName string
	Solver        func(lines AocStringArray) Any
	Elapsed       time.Duration
}

func (problem *Problem) Solve() string {
	startTime := time.Now()
	result := problem.Solver(problem.readInputToLines())
	endTime := time.Now()
	problem.Elapsed = endTime.Sub(startTime)

	switch v := result.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (problem *Problem) readInputToLines() AocStringArray {
	file, err := os.Open(problem.InputFileName)

	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines AocStringArray
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lines = append(lines, AocString(fileScanner.Text()))
	}

	err = fileScanner.Err()

	if err != nil {
		fmt.Println(err)
	}

	return lines
}
