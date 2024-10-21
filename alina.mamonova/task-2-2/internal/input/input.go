package input

import (
	"bufio"
	"fmt"
	"github.com/hahapathetic/task-2-2/internal/errors"
	"os"
	"strconv"
	"strings"
)

func checkInputBorders(num int, min, max int) bool {
	return (num >= min) && (num <= max)
}

func ProcessNumInput(intro string, min, max int) int {
	fmt.Print(intro)
	var result int = 0
	reader := bufio.NewReader(os.Stdin)
	input, err1 := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err1 == nil {
		result = checkIntInput(input, min, max)
	}
	return result
}

func checkIntInput(input string, min, max int) int {
	var result int = 0
	num, err2 := strconv.Atoi(input)
	if err2 == nil && checkInputBorders(num, min, max) {
		result = num
	} else {
		fmt.Println(errors.InvalidInput)
		os.Exit(1)
	}
	return result
}
