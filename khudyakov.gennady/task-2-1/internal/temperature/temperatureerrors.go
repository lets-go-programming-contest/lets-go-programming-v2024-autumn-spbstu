package temperature

import (
	"fmt"
	"strings"
)

func temperOutOfRangeErr(constraints TemperatureConstraints) error {
	return fmt.Errorf("некорректное значение температуры. Допускаются значения от %v до %v", constraints.Min, constraints.Max)
}

func temperTermFormatErr(validValues ...string) error {
	return fmt.Errorf("некорректное условие для температуры. Допустимы: %v", strings.Join(validValues, ", "))
}

func temperFormatErr(actual string) error {
	return fmt.Errorf("некорректный ввод. Ожидалось: <=/>= *числовое значение*. Получено: %v", actual)
}
