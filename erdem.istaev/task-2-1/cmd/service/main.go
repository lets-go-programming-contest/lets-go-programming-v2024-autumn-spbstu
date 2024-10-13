package main

import (
	"bufio"
	"fmt"
	"os"

	"erdem.istaev/task-2-1/internal/input"
	"erdem.istaev/task-2-1/internal/structure"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите количество отделов: ")
	n, err := input.ReadInt(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	bounds := make([]structure.TemperatureBounds, n)

	for i := 0; i < n; i++ {
		bounds[i] = structure.TemperatureBounds{Lower: 15, Upper: 30}
	}

	var condition string
	var k, temp int
	for i := 0; i < n; i++ {
		fmt.Print("Введите количество сотрудников: ")
		k, err = input.ReadInt(reader)
		if err != nil {
			fmt.Println(err)
		}

		for j := 0; j < k; j++ {
			fmt.Print("Укажите температуру: ")
			condition, temp, err = input.ReadCondition(reader)
			if err != nil {
				fmt.Println(err)
				return
			}

			bounds[i].UpdateBounds(condition, temp)
			if bounds[i].IsValid() {
				fmt.Println(bounds[i].Lower)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
