package main

import (
	"fmt"

	"github.com/Madyarov-Gleb/task-2-2/internal/input"
	"github.com/Madyarov-Gleb/task-2-2/internal/solving"
)

func main() {
	var (
		numb []int
		k    int
	)
	numb, k = input.InputVal()
	fmt.Println(solving.FindKMax(&numb, k))
}
