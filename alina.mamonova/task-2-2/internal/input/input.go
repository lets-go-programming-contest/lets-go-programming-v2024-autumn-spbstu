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

func ProcessNumInput(intro string, min, max int, reader *bufio.Reader) int {
	fmt.Print(intro)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var result int = 0
	if err == nil {
		result = checkIntInput(input, min, max)
	}
	return result
}

func ProcessHeapInput(min, max int, size int, reader *bufio.Reader) []int {
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
	num, err := strconv.Atoi(input)
	if err == nil && checkInputBorders(num, min, max) {
		result = num
	} else {
		fmt.Println(errors.InvalidInput)
		os.Exit(1)
	}
	return result
}
