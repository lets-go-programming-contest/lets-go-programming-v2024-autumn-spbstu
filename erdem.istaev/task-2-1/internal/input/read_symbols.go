package input

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

var (
	ErrIncorrectNumber       = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
	ErrIncorrectSeparator    = errors.New("Строка должна содержать ровно один пробел для разделения.")
	ErrIncorrectComparsionOp = errors.New("Некорректная операция сравнения. Пожалуйста, введите \">=\" или \"<=\".")
)

func ReadInt(reader *bufio.Reader) (int, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	remSpc := strings.TrimSpace(str)
	var res int
	if res, err = strconv.Atoi(remSpc); err != nil {
		return 0, ErrIncorrectNumber
	}
	return res, nil
}

func ReadCondition(reader *bufio.Reader) (string, int, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, err
	}
	remSpc := strings.TrimSpace(str)
	parts := strings.Split(remSpc, " ")
	if len(parts) != 2 {
		return "", 0, ErrIncorrectSeparator
	}

	if parts[0] != ">=" && parts[0] != "<=" {
		return "", 0, ErrIncorrectComparsionOp
	}

	var temp int
	if temp, err = strconv.Atoi(parts[1]); err != nil {
		return "", 0, ErrIncorrectNumber
	}

	return parts[0], temp, nil
}
