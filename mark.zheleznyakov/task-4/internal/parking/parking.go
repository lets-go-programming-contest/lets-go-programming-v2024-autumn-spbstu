package parking

import (
    "sync"
)

type Parking struct {
  mu sync.Mutex
  reserved int
  capacity int
}

func NewParking(capacity int) *Parking {
  return &Parking{capacity: capacity}
}

func (p *Parking) Capacity() int {
  return p.capacity
}

