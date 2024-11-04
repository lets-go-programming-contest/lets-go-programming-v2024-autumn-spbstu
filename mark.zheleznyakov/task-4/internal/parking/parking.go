package parking

import (
  "fmt"
	"sync"
)

type Parking struct {
	mu       sync.Mutex
	reserved int
	capacity int
}

func NewParking(capacity int) *Parking {
	return &Parking{capacity: capacity}
}

func (p *Parking) Capacity() int {
  p.mu.Lock()
  defer p.mu.Unlock()
	return p.capacity
}

func (p *Parking) PushCar() error {
  p.mu.Lock()
  defer p.mu.Unlock()
  if p.reserved >= p.capacity {
    return fmt.Errorf("cannot push car: parking is full")
  }
  p.reserved++
	return nil
}

func (p *Parking) PopCar() error {
  p.mu.Lock()
  defer p.mu.Unlock()
  if p.reserved <= 0 {
    return fmt.Errorf("cannot pop car: parking is empty")
  }
  p.reserved--
	return nil
}
