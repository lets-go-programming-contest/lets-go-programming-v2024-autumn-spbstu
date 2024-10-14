package input

import (
	"container/heap"
	"fmt"
	"task-2-2/internal/myErrors/inputError"
	"task-2-2/pkg/heapInt"
)

func InputData() (int, *heapInt.HeapInt, error) {

	var dishCount, grade, dishNum int
	dishes := &heapInt.HeapInt{}
	heap.Init(dishes)

	_, err := fmt.Scan(&dishCount)
	if err != nil || dishCount < 0 {
		return 0, nil, &inputError.IncorrectNumber{ErrorPlace: "count of dishes"}
	}

	for i := 0; i < dishCount; i++ {
		_, err = fmt.Scan(&grade)
		if err != nil {
			return 0, nil, &inputError.IncorrectNumber{ErrorPlace: "grade for dish"}
		}
		heap.Push(dishes, grade)
	}

	_, err = fmt.Scan(&dishNum)
	if err != nil || dishNum <= 0 || dishNum > dishCount {
		return 0, nil, &inputError.IncorrectNumber{ErrorPlace: "number of dish"}
	}

	return dishNum, dishes, nil

}
