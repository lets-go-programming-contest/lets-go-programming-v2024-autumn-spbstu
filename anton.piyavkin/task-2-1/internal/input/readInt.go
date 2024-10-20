package input

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ReadInt(in io.Reader) (int, error) {
	reader := bufio.NewReader(in)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return value, nil
}