package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	values []int
	lock   sync.RWMutex
}

func (q *Queue) Push(v int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.values = append(q.values, v)
	fmt.Printf("Added %v\n", v)
}

func (q *Queue) Pop() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.values) == 0 {
		fmt.Println("IS EMPTY")
		return 0
	}
	res := (q.values[0])
	q.values = q.values[1:]
	fmt.Printf("Removed %v\n", res)
	return res
}

func (q *Queue) PrintAll() {
	q.lock.RLock()
	defer q.lock.RUnlock()
	fmt.Printf("Current queue: %v\n", q.values)
}
