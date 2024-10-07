package input

import (
	"errors"
	"fmt"
)

func InputCheckCount() (int, error) {
	var count int
	if _, err := fmt.Scan(&count); err != nil || count <= 0 {
		return 0, errors.New("invalid entry of number of departments")
	}
	return count, nil
}

func InputCheckNewTemp() (string, float64, error) {

	const (
		commonMin = 15
		commonMax = 30
	)

	var (
		sign    string
		newTemp float64
	)

	if _, err := fmt.Scan(&sign, &newTemp); err != nil {
		return "", 0, errors.New("invalid sign or temperature input")
	}

	if newTemp < commonMin || newTemp > commonMax {
		return "", 0, errors.New("invalid temperature input - out of range (15 - 30)")
	}

	return sign, newTemp, nil
}
