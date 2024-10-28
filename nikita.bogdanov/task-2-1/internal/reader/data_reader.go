package reader

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

const (
	Greater        = ">="
	Less           = "<="
	MaxTemperature = 30
	MinTemperature = 15
)

var (
	MainDataError   = errors.New("Main data must be integer\n")
	ComparatorError = errors.New("wrong comparator format\n")
	TempFormatError = errors.New("wrong temperature's number format\n")
	TempRangeError  = errors.New("wrong temperature's number range\n")
)

func simpleRead(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.ReplaceAll(data, "\n", "")
	return data, nil
}

func ReadMainData(in io.Reader) (int, error) {
	data, err := simpleRead(in)
	if err != nil {
		return 0, err
	}
	data = strings.ReplaceAll(data, " ", "")
	result, err := strconv.Atoi(data)
	if err != nil {
		return 0, MainDataError
	}
	return result, nil
}

func ReadConditionerData(in io.Reader) (int, string, error) {
	data, err := simpleRead(in)
	if err != nil {
		return 0, "", err
	}
	data = strings.ReplaceAll(data, " ", "")
	comp := string(data[:2])
	num := data[2:]
	if comp != Greater && comp != Less {
		return 0, "", ComparatorError
	}
	number, err := strconv.Atoi(num)
	if err != nil {
		return 0, "", TempFormatError
	}
	if number > MaxTemperature || number < MinTemperature {
		return 0, comp, TempRangeError
	}
	return number, comp, nil
}
