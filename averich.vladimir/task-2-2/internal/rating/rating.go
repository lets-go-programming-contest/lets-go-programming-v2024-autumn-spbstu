package rating

import (
	"fmt"
	"sort"
	"task-2-2/internal/heapOfDishes"
	"task-2-2/internal/userErrors"
)

func DetectTheMostGuessedDish(index int, heapOfDishes *heapOfDishes.HeapOfDishes) (int, error) {

	sort.Sort(heapOfDishes)

	if heapOfDishes.Len() < index+1 {
		return 0, fmt.Errorf("%w", userErrors.ErrorOutOfRange)
	}

	var result int
	for i := 0; i < index; i++ {
		result = heapOfDishes.Pop().(int)
	}

	return result, nil
}
