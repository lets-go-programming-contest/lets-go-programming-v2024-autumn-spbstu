package entity

import (
	"errors"
	"fmt"
	"sync"
)

type UnsafeStore struct {
	Products map[string]int
}

func (s *UnsafeStore) Add(name string, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	if _, ok := s.Products[name]; ok {
		s.Products[name] += count
	} else {
		s.Products[name] = count
	}
	fmt.Printf("Product %s added to store. Count = %d\n", name, s.Products[name])
}

func (s *UnsafeStore) Sell(name string, sellCount int, wg *sync.WaitGroup) error {
	defer wg.Done()
	productCount, ok := s.Products[name]
	if !ok {
		return errors.New("product not found")
	}

	if productCount-sellCount < 0 {
		return fmt.Errorf("product %s not enough to sell %d. Count = %d", name, sellCount, productCount)
	}

	s.Products[name] -= sellCount
	fmt.Printf("Product %s sold from store. Count = %d\n", name, s.Products[name])
	return nil
}
