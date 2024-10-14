package input

import (
	"bufio"
	"errors"
	"strconv"
	"strings"

	"erdem.istaev/task-1/internal/calculator"
)

var (
	ErrIncorrectNumber = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
)

func ReadFloat64(reader *bufio.Reader) (float64, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	remSpc := strings.TrimSpace(str)
	var res float64
	if res, err = strconv.ParseFloat(remSpc, 64); err != nil {
		return 0, ErrIncorrectNumber
	}
	return res, nil
}

func ReadOperation(reader *bufio.Reader) (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	operation := strings.TrimSpace(str)
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		return "", calculator.ErrIncorrectOperation
	}
	return operation, nil
}
