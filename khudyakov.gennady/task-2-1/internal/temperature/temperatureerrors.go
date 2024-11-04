package temperature

import (
	"fmt"
)

type TemperatureTermFormatError struct {
	Actual string
}

func (err TemperatureTermFormatError) Error() string {
	return fmt.Sprintf("некорректное условие для температуры. Допустимы: <= или >=, получено: %v", err.Actual)
}

type TemperatureFormatError struct {
	Actual string
}

func (err TemperatureFormatError) Error() string {
	return fmt.Sprintf("некорректный ввод. Ожидалось: <=/>= *числовое значение*. Получено: %v", err.Actual)
}
