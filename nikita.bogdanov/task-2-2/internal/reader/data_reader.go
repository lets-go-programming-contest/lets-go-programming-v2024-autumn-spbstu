package reader

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	MainDataError    = errors.New("Main data must be integer\n")
	MealsNumberError = errors.New("Wrong number of meals\n")
	MealNumberError  = errors.New("Wrong number of meal\n")
)

const (
	UpMealBound   = 10000
	DownMealBound = -10000
)

func baseRead(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.ReplaceAll(data, "\n", "")
	return data, nil
}

func IntMainDataRead(in io.Reader) (int, error) {
	data, err := baseRead(in)
	if err != nil {
		return 0, err
	}
	data = strings.ReplaceAll(data, " ", "")
	result, err := strconv.Atoi(data)
	if err != nil {
		return 0, MainDataError
	}
	return result, nil
}

func IntMealRead(in io.Reader, n int) ([]int, error) {
	meals := make([]int, n)
	mealsTmp, err := baseRead(in)
	if err != nil {
		return nil, err
	}
	mealsTmpStr := strings.Fields(mealsTmp)
	if len(mealsTmpStr) != n {
		return nil, MealsNumberError
	}
	for idx, dish := range mealsTmpStr {
		meal, err := strconv.Atoi(dish)
		if err != nil {
			return nil, err
		}
		if meal < DownMealBound || meal > UpMealBound {
			return nil, MealNumberError
		}
		meals[idx] = meal
	}
	return meals, nil
}
