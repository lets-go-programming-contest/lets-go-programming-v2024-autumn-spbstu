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

func ProcessNumInput(intro string) int {
	fmt.Print(intro)
	var result int = 0
	reader := bufio.NewReader(os.Stdin)
	input, err1 := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err1 == nil {
		result = checkIntInput(input)
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

func checkIntInput(input string) int {
	var result int = 0
	num, err2 := strconv.Atoi(input)
	if err2 == nil && checkInputBorders(num, minInputBound, maxInputBound) {
		result = num
	} else {
		fmt.Println(errors.InvalidInput)
		os.Exit(1)
	}
	return result
}

func ProcessFullInput() (string, int) {
	var firstPart string = ""
	var secondPart int = 0
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputParts := strings.Fields(input)

	if len(inputParts) != 2 {
		fmt.Println(errors.InvalidInput)
		os.Exit(1)
	}

	if err == nil {
		firstPart = checkOperatorInput(inputParts[0])
		secondPart = checkIntInput(inputParts[1])
	}
	return firstPart, secondPart
}
