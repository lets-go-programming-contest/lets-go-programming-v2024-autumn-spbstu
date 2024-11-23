package readers

import (
	"fmt"
)

func ReadIntNum(message string, min int, max int) int {
	var num float64
	var resultNum int

	if message != "" {
		fmt.Print(message)
	}

	_, err := fmt.Scan(&num)
	if err != nil {
		panic("ошибка: некорректное значение")
	}

	resultNum = int(num)
	if num != float64(resultNum) {
		panic("ошибка: число должно быть целым")
	}

	if resultNum < min || resultNum > max {
		panic(fmt.Sprintf("ошибка: значение должно быть в диапазоне от %d до %d", min, max))
	}

	return resultNum
}
