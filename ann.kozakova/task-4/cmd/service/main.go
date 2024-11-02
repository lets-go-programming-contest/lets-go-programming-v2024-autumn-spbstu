package main

import (
	"fmt"
	"github.com/nutochk/task-4/internal/shops"
	"sync"
)

func main() {
	shop := shops.Shop{Items: make([]string, 0)}
	shop.Items = append(shop.Items, "Product0", "Product1")
	shop.All = 2
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			all := shop.All
			shops.Add100Items(&shop, all)
			fmt.Println("Supplier", i, "was", all, "now", shop.All)
		}(i)
	}
	wg.Wait()
}
