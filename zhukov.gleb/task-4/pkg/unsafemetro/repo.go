package metro

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func UnsafeNewVisitorsRepo() *UnsafeVisitors {
	return &UnsafeVisitors{
		Quantity: make(map[int]struct{}, 10),
	}
}

func (v *UnsafeVisitors) UnsafeRegister(id int) {
	v.Quantity[id] = struct{}{}
}

func (v *UnsafeVisitors) UnsafeGetCount() int {
	return len(v.Quantity)
}

func (v *UnsafeVisitors) UnsafeSimulator(cntVisitors int, out chan<- string) {
	wg := &sync.WaitGroup{}
	wg.Add(cntVisitors)

	go func() {
		for {
			time.Sleep(time.Millisecond * 200)

			count := v.UnsafeGetCount()

			out <- fmt.Sprintf("count visitors: %d", count)
		}
	}()

	for i := 0; i < cntVisitors; i++ {
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))
			v.UnsafeRegister(i)

			out <- fmt.Sprintf("added user: %d", i)
		}(i)
	}

	wg.Wait()
	out <- "program simulation ended successfully"
	close(out)
}
