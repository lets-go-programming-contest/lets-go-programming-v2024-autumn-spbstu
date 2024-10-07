package input

import (
	"errors"
	"fmt"
)

func InputCheckCount() (int, error) {
	var count int
	if _, err := fmt.Scan(&count); err != nil || count <= 0 {
		return 0, errors.New("неверный ввод количества отделов")
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
		return "", 0, errors.New("неверный ввод знака или температуры")
	}

	if newTemp < commonMin || newTemp > commonMax {
		return "", 0, errors.New("неверный ввод температуры - выход за допустимый диапазон (15 - 30)")
	}

	return sign, newTemp, nil
}
