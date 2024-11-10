package input

import (
	"bufio"
	"fmt"
	"github.com/hahapathetic/task-2-1/internal/errors"
	"os"
	"strconv"
	"strings"
)

const (
	minInputBound = 1
	maxInputBound = 1000
)

func checkInputBorders(num int, min, max int) bool {
	return (num >= min) && (num <= max)
}

func ProcessNumInput(intro string, reader *bufio.Reader) int {
	fmt.Print(intro)
	var result int = 0
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err == nil {
		result = checkIntInput(input, true)
	}
	return result
}

func checkOperatorInput(input string) string {
	var result string = ""
	if input == "<=" || input == ">=" {
		result = input
	} else {
		fmt.Println(errors.InvalidOperator)
		os.Exit(1)
	}
	return result
}

func checkIntInput(input string, border bool) int {
	var result int = 0
	num, err := strconv.Atoi(input)
	if err == nil {
		if border {
			checkInputBorders(num, minInputBound, maxInputBound)
		}
		result = num
	} else {
		fmt.Println(errors.InvalidInput)
		os.Exit(1)
	}
	return result
}

func ProcessFullInput(reader *bufio.Reader) (string, int) {
	var firstPart string = ""
	var secondPart int = 0
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputParts := strings.Fields(input)

	if len(inputParts) != 2 {
		fmt.Println(errors.InvalidInput)
		os.Exit(1)
	}

	if err == nil {
		firstPart = checkOperatorInput(inputParts[0])
		secondPart = checkIntInput(inputParts[1], false)
	}
	return firstPart, secondPart
}
