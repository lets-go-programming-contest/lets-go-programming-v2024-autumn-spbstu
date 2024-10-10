package main

import (
	"fmt"
	"log"

	"github.com/nutochk/task-2-2/internal/data"
	"github.com/nutochk/task-2-2/internal/max"
)

func main() {

	numbers, k := data.EnterData()

	kMax, err := max.FindKMax(&numbers, k)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kMax)
}
