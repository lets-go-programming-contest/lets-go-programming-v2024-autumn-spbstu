package parking

import (
	"sync"
)

type Parking struct {
	mu       sync.Mutex
	in int
	out int
}

func NewParking() *Parking {
	return &Parking{}
}

func (p *Parking) Popped() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.out
}

func (p *Parking) Pushed() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.in
}

func (p *Parking) PushCar() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.in++
	return nil
}

func (p *Parking) PopCar() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.out++
	return nil
}

func (p *Parking) UnsafePushCar() error {
	p.in++
	return nil
}

func (p *Parking) UnsafePopCar() error {
	p.out++
	return nil
}
