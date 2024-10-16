package io

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadInt(reader *bufio.Reader) (int, error) {
	inputString, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseInt(strings.TrimSpace(inputString), 10, 0)
	if err != nil {
		return 0, ParseIntError(inputString)
	}

	return int(result), nil
}
