package reader

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

const (
	greater = ">="
	less    = "<="
)

func simple_read(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.ReplaceAll(data, "\n", "")
	return data, nil
}

func Read_main_data(in io.Reader) (int, error) {
	data, err := simple_read(in)
	if err != nil {
		return 0, err
	}
	data = strings.ReplaceAll(data, " ", "")
	result, err := strconv.Atoi(data)
	if err != nil {
		return 0, errors.New("Main data must be integer")
	}
	return result, nil
}

func Read_conditioner_data(in io.Reader) (int, string, error) {
	data, err := simple_read(in)
	if err != nil {
		return 0, "", err
	}
	data = strings.ReplaceAll(data, " ", "")
	comp := string(data[:2])
	num := data[2:]
	if comp != greater && comp != less {
		return 0, "", errors.New("wrong comparator format\n")
	}
	number, err := strconv.Atoi(num)
	if err != nil {
		return 0, "", errors.New("wrong temperature's number format\n")
	}
	if number > 30 || number < 15 {
		return 0, comp, errors.New("wrong temperature's number range\n")
	}
	return number, comp, nil
}
