package main

import (
	"fmt"

	"github.com/Mmmakskl/task-2-2/internal/input"
	"github.com/Mmmakskl/task-2-2/internal/logic"
)

func main() {
	k, rating := input.ReadNumber()

	fmt.Printf("Result: %d\n", logic.KMax(k, rating))
}
