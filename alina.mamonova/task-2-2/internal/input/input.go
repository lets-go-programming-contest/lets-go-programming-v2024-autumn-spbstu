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
	reader := bufio.NewReader(os.Stdin)
	input, err1 := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var result int = 0
	if err1 == nil {
		result = checkIntInput(input, min, max)
	}
	return result
}

func ProcessHeapInput(min, max int, size int) []int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil
	}

	input = strings.TrimSpace(input)
	values := strings.Fields(input)

	var results []int
	for _, val := range values {
		num, err := strconv.Atoi(val)
		if err != nil || num < min || num > max {
			fmt.Printf("Invalid input.")
			os.Exit(1)
		}
		results = append(results, num)
	}
	if len(values) != size {
		fmt.Printf("Invalid input.")
		os.Exit(1)
	}

	return results
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
