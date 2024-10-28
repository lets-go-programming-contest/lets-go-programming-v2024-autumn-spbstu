package main

import (
	"fmt"
	"log"

	"github.com/Mmmakskl/task-2-2/internal/input"
	"github.com/Mmmakskl/task-2-2/internal/logic"
)

func main() {
	k, rating, err := input.ReadNumber()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", logic.KMax(k, rating))
}
