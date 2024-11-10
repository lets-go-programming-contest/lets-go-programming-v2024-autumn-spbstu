package input

import (
	"container/heap"
	"fmt"

	"github.com/Mmmakskl/task-2-2/internal/structure"
)

func ReadNumber() (int, *structure.IntHeap, error) {
	var (
		n  int
		k  int
		ai int
	)

	fmt.Print("Enter the number of dishes: ")
	rating := &structure.IntHeap{}
	_, err := fmt.Scanln(&n)
	if err != nil || n <= 0 {
		return 0, rating, ErrInput
	}

	fmt.Print("Enter the rating dishes: ")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&ai)
		if err != nil {
			return 0, rating, ErrInput
		}
		heap.Push(rating, ai)
	}

	fmt.Print("Enter the sequence number of the k-th dish: ")
	_, err = fmt.Scanln(&k)
	if err != nil || k <= 0 {
		return 0, rating, ErrInput
	}
	if rating.Len() < k {
		return 0, rating, ErrK_th
	}

	return k, rating, nil
}
