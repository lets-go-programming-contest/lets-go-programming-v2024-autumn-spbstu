package io

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrorIncorectInt       = errors.New("> Некорректное значение запускаемого теста. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectIntBounds = errors.New("> Некорректное количество. Значение должно быть равным либо 1, либо 2!")
)

func GetNumber(prompt string, minVal int, maxVal int, reader *bufio.Reader) (int, error) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
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
