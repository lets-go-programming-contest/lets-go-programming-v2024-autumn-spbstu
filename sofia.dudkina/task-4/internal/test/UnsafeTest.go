package test

import (
	"fmt"
	"sync"

	"github.com/sssidkn/task-4/internal/entity"
)

func UnsafeTest() {
	wg := &sync.WaitGroup{}
	store := &entity.UnsafeStore{}
	store.Products = map[string]int{"Milk": 10, "Apple": 2, "Orange": 3}
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func() {
			err := store.Sell("Orange", i, wg)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go store.Add("Orange", i, wg)
	}
	wg.Wait()
}
