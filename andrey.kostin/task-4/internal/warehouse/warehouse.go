package warehouse

import (
	"sync"
)

type Warehouse struct {
	items   map[string]int
	rwMutex sync.RWMutex
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		items: make(map[string]int),
	}
}
