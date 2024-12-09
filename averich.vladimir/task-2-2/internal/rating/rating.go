package rating

import (
	"sort"
	"task-2-2/internal/heapOfDishes"
	"task-2-2/internal/userErrors"
)

func DetectTheMostGuessedDish(index int, heapOfDishes *heapOfDishes.HeapOfDishes) (int, error) {
	sort.Sort(sort.Reverse(heapOfDishes))
	if value, ok := heapOfDishes.Get(index - 1); ok {
		return value, nil
	} else {
		return 0, userErrors.ErrorInvalidValue
	}
}
