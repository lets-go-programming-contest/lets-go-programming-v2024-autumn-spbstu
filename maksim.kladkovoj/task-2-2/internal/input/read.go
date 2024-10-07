package input

import (
	"container/heap"
	"errors"
	"fmt"
	"log"

	"github.com/Mmmakskl/task-2-2/internal/structure"
)

func ReadNumber() (int, *structure.IntHeap) {
	var (
		n        int
		k        int
		ai       int
		errInput error = errors.New("Input error")
		errK_th  error = errors.New("The number of dishes is less than the k-th number")
	)

	fmt.Print("Enter the number of dishes: ")
	rating := &structure.IntHeap{}
	_, err := fmt.Scanln(&n)
	if err != nil || n <= 0 {
		log.Fatal(errInput)
	}

	fmt.Print("Enter the rating dishes: ")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&ai)
		if err != nil {
			log.Fatal(errInput)
		}
		heap.Push(rating, ai)
	}

	fmt.Print("Enter the sequence number of the k-th dish: ")
	_, err = fmt.Scanln(&k)
	if err != nil || k <= 0 {
		log.Fatal(errInput)
	}
	if rating.Len() < k {
		log.Fatal(errK_th)
	}

	return k, rating
}
