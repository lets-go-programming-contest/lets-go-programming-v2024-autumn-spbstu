package main

import (
	"fmt"

	"erdem.istaev/task-2-2/internal/findkmax"
	"erdem.istaev/task-2-2/internal/input"
)

func main() {
	k, dishes, n, err := input.ReadData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(findkmax.FindKMax(k, dishes, n))
}
