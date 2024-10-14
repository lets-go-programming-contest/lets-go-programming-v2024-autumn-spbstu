package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() ([]int, int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 || n > 10000 {
		return nil, 0, fmt.Errorf("Ошибка: некорректный ввод количества блюд ")
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	mealsStr := strings.Fields(input)

	if len(mealsStr) != n {
		return nil, 0, fmt.Errorf("Ошибка: некорректный ввод последовательности ")
	}

	meals := make([]int, n)
	for i, mealStr := range mealsStr {
		meal, err := strconv.Atoi(mealStr)
		if err != nil || meal < -10000 || meal > 10000 {
			return nil, 0, fmt.Errorf("Ошибка: некорректный ввод последовательности ")
		}
		meals[i] = meal
	}

	var k int
	_, err = fmt.Scanln(&k)
	if err != nil || k < 1 || k > n {
		return nil, 0, fmt.Errorf("Ошибка: некорректный ввод порядкового номера k-го по предпочтению блюда ")
	}

	return meals, k, nil
}
