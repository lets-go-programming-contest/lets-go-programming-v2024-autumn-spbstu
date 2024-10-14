package main

import (
	"anastasiya.soboleva/task-2-1/internal/temperature"
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 || n > 1000 {
		log.Fatal("Ошибка: некорректное количество отделов")
	}

	for i := 0; i < n; i++ {
		var k int
		_, err := fmt.Scanln(&k)
		if err != nil || k < 1 || k > 1000 {
			log.Fatal("Ошибка: некорректное количество сотрудников в отделе")
		}

		tracker := temperature.NewTracker()

		for j := 0; j < k; j++ {
			var operation string
			var value int
			_, err := fmt.Scanf("%s %d\n", &operation, &value)
			if err != nil {
				log.Fatal("Ошибка: некорректный ввод данных")
			}
			result, err := tracker.ProcessTemperature(operation, value)
			if err != nil {
				log.Fatal("Ошибка:", err)
			}
			fmt.Println(result)
		}
	}
}
