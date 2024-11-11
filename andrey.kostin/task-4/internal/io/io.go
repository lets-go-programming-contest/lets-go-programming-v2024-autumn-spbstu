package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrorIncorectIntBounds = errors.New("> Некорректное значение запускаемого теста. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectInt       = errors.New("> Некорректное количество. Значение должно быть равным либо 1, либо 2!")
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
			return 0, ErrorIncorectIntBounds
		}
	} else {
		return 0, ErrorIncorectInt
	}
}
