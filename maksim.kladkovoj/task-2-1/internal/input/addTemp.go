package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const expectedMatchCount = 3

type OperatorType int

const (
	Unknown OperatorType = iota
	Less
	LessOrEqual
	Greater
	GreaterOrEqual
)

func ParseTemp(input string) (OperatorType, int, error) {
	re, err := regexp.Compile(`([<>]=?)\s*(\d+)`)
	if err != nil {
		return Unknown, 0, ErrRegexp
	}

	matches := re.FindStringSubmatch(input)
	if len(matches) < expectedMatchCount {
		return Unknown, 0, ErrInput
	}

	operator := matches[1]
	var operatorType OperatorType
	switch operator {
	case "<":
		operatorType = Less
	case "<=":
		operatorType = LessOrEqual
	case ">":
		operatorType = Greater
	case ">=":
		operatorType = GreaterOrEqual
	default:
		operatorType = Unknown
	}

	temp, err := strconv.Atoi(matches[2])
	if err != nil {
		return Unknown, 0, ErrInput
	}

	return operatorType, temp, nil
}

func AddNumber() (int, error) {
	var number int

	_, err := fmt.Scanln(&number)
	if err != nil {
		return 0, ErrInput
	}

	return number, nil
}

func AddTemperature() (OperatorType, int, error) {
	var (
		operator OperatorType
	)
	in := bufio.NewReader(os.Stdin)

	value, err := in.ReadString('\n')
	if err != nil {
		return Unknown, 0, ErrInput
	}

	operator, temperature, err := ParseTemp(value)
	if err != nil {
		return Unknown, 0, ErrTemp
	}

	return operator, temperature, nil
}
