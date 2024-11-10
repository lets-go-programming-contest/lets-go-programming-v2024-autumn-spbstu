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

	minTemp := 15
	maxTemp := 30

	for i := 0; i < n; i++ {
		bounds[i] = structure.TemperatureBounds{Lower: minTemp, Upper: maxTemp}
	}

	input.SetTemperature(reader, n, bounds)
}
