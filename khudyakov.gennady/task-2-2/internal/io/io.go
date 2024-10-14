package io

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ParseIntError(actual string) error {
	return fmt.Errorf("ожидалось целое число. Получено: %v", actual)
}

func ReadInt(reader *bufio.Reader) (int64, error) {
	inputString, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseInt(strings.TrimSpace(inputString), 10, 64)
	if err != nil {
		return 0, ParseIntError(inputString)
	}

	return result, nil
}
