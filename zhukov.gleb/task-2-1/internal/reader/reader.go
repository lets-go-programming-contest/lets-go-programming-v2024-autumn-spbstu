package reader

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	minNK   = 0
	maxNK   = 1000
	minTemp = 15
	maxTemp = 30
)

var (
	ErrInvalidNK   = errors.New("плохое значение NK")
	ErrInvalidTemp = errors.New("плохое значение температуры")
)

func readData(number *int) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	input = strings.TrimSpace(input)
	if input == "выход" {
		return true, nil
	}
	*number, err = strconv.Atoi(input)
	return false, err
}

func ReadNK() (int, bool, error) {
	var cntPlaces int
	exit, err := readData(&cntPlaces)
	if exit || err != nil {
		return 0, exit, err
	}
	if cntPlaces <= minNK || cntPlaces > maxNK {
		return 0, false, ErrInvalidNK
	}
	return cntPlaces, false, nil
}

func ReadCondition() (string, bool, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", false, err
	}
	input = strings.TrimSpace(input)
	if input == "выход" {
		return "", true, nil
	}
	return input, false, nil
}

func ParseTemperature(input string) (int, error) {
	temp, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrInvalidTemp
	}
	if temp < minTemp || temp > maxTemp {
		return 0, ErrInvalidTemp
	}
	return temp, nil
}
