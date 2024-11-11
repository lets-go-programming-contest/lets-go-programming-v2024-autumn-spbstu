package test

import (
	"fmt"
	"sync"
	"time"

	"github.com/IDevFrye/task-4/internal/warehouse"
)

func SafeTest() {
	wh := warehouse.NewWarehouse()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(step int) {
			defer wg.Done()
			wh.AddItem("itemA", 1)
			fmt.Printf("Шаг %d: Добавили 1, текущее количество itemA (безопасно): %d\n", step, wh.GetQuantity("itemA"))
		}(i)

		go func(step int) {
			defer wg.Done()
			wh.RemoveItem("itemA", 1)
			fmt.Printf("Шаг %d: Убрали 1, текущее количество itemA (безопасно): %d\n", step, wh.GetQuantity("itemA"))
		}(i)
	}

	wg.Wait()
	fmt.Println("Итоговое количество itemA (безопасный режим):", wh.GetQuantity("itemA"))
}

func UnsafeTest() {
	wh := warehouse.NewWarehouse()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(step int) {
			defer wg.Done()
			wh.AddItemUnsafe("itemA", 1)
			fmt.Printf("Шаг %d: Добавили 1, текущее количество itemA (небезопасно): %d\n", step, wh.GetQuantityUnsafe("itemA"))
		}(i)

		go func(step int) {
			defer wg.Done()
			wh.RemoveItemUnsafe("itemA", 1)
			fmt.Printf("Шаг %d: Убрали 1, текущее количество itemA (небезопасно): %d\n", step, wh.GetQuantityUnsafe("itemA"))
		}(i)

		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
	fmt.Println("Итоговое количество itemA (небезопасный режим):", wh.GetQuantityUnsafe("itemA"))
}
