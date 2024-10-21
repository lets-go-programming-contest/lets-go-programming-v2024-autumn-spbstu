package reader

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
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
		return 0, errors.New("Main data must be integer")
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
		return nil, errors.New("Wrong number of meals")
	}
	for idx, dish := range mealsTmpStr {
		meal, err := strconv.Atoi(dish)
		if err != nil {
			return nil, err
		}
		if meal < -10000 || meal > 10000 {
			return nil, errors.New("Wrong number of meal\n")
		}
		meals[idx] = meal
	}
	return meals, nil
}
