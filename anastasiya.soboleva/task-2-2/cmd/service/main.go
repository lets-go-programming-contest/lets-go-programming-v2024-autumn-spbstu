package main

import (
	"anastasiya.soboleva/task-2-1/internal/heap"
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 || n > 10000 {
		log.Fatal("Ошибка: неккоректный ввод количества блюд")
	}

	meals := make([]int, n)
	for i := 0; i < n; i++ {
		var meal int
		_, err := fmt.Scan(&meal)
		if err != nil || meal < -10000 || meal > 10000 {
			log.Fatal("Ошибка: неккоректный ввод последовательности")
		}
		meals[i] = meal
	}

	var k int
	_, err = fmt.Scanln(&k)
	if err != nil || k < 1 || k > n {
		log.Fatal("Ошибка: неверный ввод порядкового номера k-го по предпочтению блюда")
	}

	result := heap.FindKMeal(meals, k)

	fmt.Printf("%d", result)
}
