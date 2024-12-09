package input

import (
	"fmt"
	"task-2-1/internal/userErrors"
)

type DiapasonOfTemperature struct {
	Lower int
	Upper int
}

func ReadQuantityOfDepts() (int, error) {

	var number int

	_, err := fmt.Scanln(&number)
	if err != nil || number < 1 || number > 1000 {
		return 0, fmt.Errorf("%w: %w", userErrors.ErrOverflow, err)
	}

	return number, nil
}

func ReadQuantityOfWorkers() (int, error) {

	var number int

	_, err := fmt.Scanln(&number)
	if err != nil || number < 1 || number > 1000 {
		return 0, fmt.Errorf("%w: %w", userErrors.ErrOverflow, err)
	}

	return number, nil
}

func ReadTemperature(rangeOfTemp DiapasonOfTemperature) (DiapasonOfTemperature, error) {

	var temperature int
	var operator string

	_, err := fmt.Scanf("%s %d\n", &operator, &temperature)
	if err != nil || temperature > rangeOfTemp.Upper || temperature < rangeOfTemp.Lower || (operator != ">=" && operator != "<=") {
		return DiapasonOfTemperature{}, fmt.Errorf("%w: %w", userErrors.ErrIncorrectInput, err)
	}

	if operator == ">=" {
		rangeOfTemp.Lower = temperature
	} else {
		rangeOfTemp.Upper = temperature
	}

	return rangeOfTemp, nil
}
