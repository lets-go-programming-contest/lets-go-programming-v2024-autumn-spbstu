package io

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

const (
	parseIntError = "ожидалось целое число"
)

func ReadInt(reader *bufio.Reader) (int64, error) {
	inputString, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseInt(strings.TrimSpace(inputString), 10, 64)
	if err != nil {
		return 0, errors.New(parseIntError)
	}

	return result, nil
}

func ReadTerm(reader *bufio.Reader) (string, int64, error) {
	inputString, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, err
	}

	split := strings.Split(strings.TrimSpace(inputString), " ")
	if len(split) != 2 {
		return "", 0, errors.New("некорректный ввод. Ожидалось: <=/>= *числовое значение*")
	}
	temperature, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return "", 0, errors.New(parseIntError)
	}
	return split[0], temperature, nil
}
