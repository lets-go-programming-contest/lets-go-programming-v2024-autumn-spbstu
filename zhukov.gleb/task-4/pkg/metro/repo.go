package metro

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func NewVisitorsRepo() *Visitors {
	return &Visitors{
		Quantity: make(map[int]struct{}, 10),
		mu:       &sync.RWMutex{},
	}
}

func (v *Visitors) Register(id int) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.Quantity[id] = struct{}{}
}

func (v *Visitors) GetCount() int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return len(v.Quantity)
}

func (v *Visitors) Simulator(cntVisitors int, out chan<- string) {
	wg := &sync.WaitGroup{}
	wg.Add(cntVisitors)

	go func() {
		for {
			time.Sleep(time.Millisecond * 200)

			count := v.GetCount()

			out <- fmt.Sprintf("count visitors: %d", count)
		}
	}()

	for i := 0; i < cntVisitors; i++ {
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))
			v.Register(i)

			out <- fmt.Sprintf("added user: %d", i)
		}(i)
	}

	wg.Wait()
	out <- "program simulation ended successfully"
	close(out)
}
