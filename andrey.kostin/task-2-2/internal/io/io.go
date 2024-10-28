package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	errProc "github.com/IDevFrye/task-2-2/internal/errors"
)

func GetNumber(prompt string, minVal int, maxVal int) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	if numb, err := strconv.Atoi(input); err == nil {
		if numb <= maxVal && numb >= minVal {
			return numb, nil
		} else {
			return 0, errProc.ErrorIncorectIntBounds
		}
	} else {
		return 0, errProc.ErrorIncorectInt
	}
}

func GetHeapElements(prompt string, minVal int, maxVal int, count int) ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	prefs := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Print(prompt)
		scanner.Scan()
		inputAi := scanner.Text()
		if numAi, err := strconv.Atoi(inputAi); err == nil {
			if numAi <= maxVal && numAi >= minVal {
				prefs[i] = numAi
			} else {
				return nil, errProc.ErrorIncorectHeapBounds
			}
		} else {
			return nil, errProc.ErrorIncorectHeap
		}
	}
	return prefs, nil
}

func GetPrefDish(prompt string, minVal int, count int) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	if numb, err := strconv.Atoi(input); err == nil {
		if numb <= count && numb >= minVal {
			return numb, nil
		} else {
			return 0, errProc.ErrorIncorectPrefDishBounds
		}
	} else {
		return 0, errProc.ErrorIncorectPrefDish
	}
}
