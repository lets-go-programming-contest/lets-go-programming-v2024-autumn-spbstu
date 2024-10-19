package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	errProc "github.com/IDevFrye/task-2-2/internal/errors"
)

func getNumber(prompt string, minVal int, maxVal int) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	if numb, err := strconv.Atoi(input); err != nil {
		if numb <= maxVal && numb >= minVal {
			return numb, nil
		} else {
			return 0, errProc.ErrorIncorectIntBounds
		}
	} else {
		return 0, errProc.ErrorIncorectInt
	}
}

func getHeapElement(prompt string, minVal int, maxVal int, count int) ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	prefs := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Print(prompt)
		scanner.Scan()
		input_i := scanner.Text()
		if numb_i, err := strconv.Atoi(input_i); err != nil {
			if numb_i <= maxVal && numb_i >= minVal {
				prefs[i] = numb_i
			} else {
				return nil, errProc.ErrorIncorectHeapBounds
			}
		} else {
			return nil, errProc.ErrorIncorectHeap
		}
	}
	return prefs, nil
}

func getPrefDish(prompt string, minVal int, count int) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	input := scanner.Text()
	if numb, err := strconv.Atoi(input); err != nil {
		if numb <= count && numb >= minVal {
			return numb, nil
		} else {
			return 0, errProc.ErrorIncorectPrefDishBounds
		}
	} else {
		return 0, errProc.ErrorIncorectPrefDish
	}
}
