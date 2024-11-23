package counter

import (
	"sync"
)

type Counter struct {
	countEnter int
	countExit  int
	mutex      sync.Mutex
}

func (counter *Counter) Enter(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.countEnter++
}

func (counter *Counter) EnterUnsafe(i int, channel chan int) {
	counter.countEnter++
	channel <- i
}

func (counter *Counter) Exit(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.countExit++
}

func (counter *Counter) ExitUnsafe(i int, channel chan int) {
	counter.countExit++
	channel <- i
}

func (counter *Counter) GetValue() (int, int) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.countEnter, counter.countExit
}
