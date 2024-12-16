package io

import (
	"container/heap"
	"fmt"
	"strconv"
	"task-2-2/internal/heapOfDishes"
	"task-2-2/internal/userErrors"
)

func inputInts() (int, error) {

	var number string

	if _, err := fmt.Scan(&number); err != nil {
		return 0, userErrors.ErrorInvalidInput
	}

	parsedNumber, err := strconv.Atoi(number)

	if err != nil {
		return 0, userErrors.ErrorInvalidInput
	}

	if parsedNumber < 1 || parsedNumber > 10000 {
		return 0, userErrors.ErrorOverflow
	}

	return parsedNumber, nil
}

func InputParameters() (int, *heapOfDishes.HeapOfDishes, error) {

	var number string

	countOfDishes, err1 := inputInts()
	if err1 != nil {
		return 0, nil, err1
	}

	dishes := &heapOfDishes.HeapOfDishes{}
	heap.Init(dishes)

	for i := 0; i < countOfDishes; i++ {

		if _, err := fmt.Scan(&number); err != nil {
			return 0, nil, err
		}

		parsedNumber, err := strconv.Atoi(number)

		if err != nil {
			return 0, nil, userErrors.ErrorInvalidInput
		}

		if parsedNumber > 10000 && parsedNumber < -10000 {
			return 0, nil, userErrors.ErrorInvalidValue
		}

		heap.Push(dishes, parsedNumber)

	}

	index, err2 := inputInts()
	if err2 != nil {
		return 0, nil, err2
	}

	return index, dishes, nil
}
