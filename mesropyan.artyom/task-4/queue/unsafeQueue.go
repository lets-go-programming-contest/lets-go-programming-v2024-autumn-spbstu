package queue

import (
	"fmt"
)

type UnsafeQueue struct {
	values []int
}

func (q *UnsafeQueue) Push(v int) {
	q.values = append(q.values, v)
	fmt.Printf("Added %v\n", v)
}

func (q *UnsafeQueue) Pop() int {
	if len(q.values) == 0 {
		fmt.Println("IS EMPTY")
		return 0
	}
	res := (q.values[0])
	q.values = q.values[1:]
	fmt.Printf("Removed %v\n", res)
	return res
}

func (q *UnsafeQueue) PrintAll() {
	fmt.Printf("Current queue: %v\n", q.values)
}
