package main

import (
	"fmt"
	"sync"

	"github.com/sssidkn/task-4/internal/entity"
)

func main() {
	store := &entity.Store{}
	wg := &sync.WaitGroup{}
	store.Products = map[string]int{"Milk": 10, "Apple": 2, "Orange": 3}
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func() {
			err := store.SellProduct("Orange", i, wg)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go store.AddProduct("Orange", i, wg)
	}
	wg.Wait()
}
