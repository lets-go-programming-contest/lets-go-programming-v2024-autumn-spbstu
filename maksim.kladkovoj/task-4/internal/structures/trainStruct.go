package structures

import (
	"sync"
)

type TrainCounter struct {
	Arrival  int
	Departed int
	mutex    sync.Mutex
}

func (counter *TrainCounter) Arrive(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.Arrival++
}

func (counter *TrainCounter) ArriveUnsafe() {
	counter.Arrival++
}

func (counter *TrainCounter) Depart(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.Departed++
}

func (counter *TrainCounter) DepartUnsafe() {
	counter.Departed++
}

func (counter *TrainCounter) GetArrival() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.Arrival
}

func (counter *TrainCounter) GetDeparted() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.Departed
}
