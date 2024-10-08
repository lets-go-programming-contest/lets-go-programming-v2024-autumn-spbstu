package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/nutochk/task-2-2/internal/max"
)

func main() {
	fmt.Println("Enter n:")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(errors.New("incorrect data"))
	}

	fmt.Println("Enter ai:")

	numbers := make([]int, n)
	i := 0
	var ai int
	for ; i < n; i++ {
		_, err = fmt.Scan(&ai)
		if err != nil {
			log.Fatal(errors.New("incorrect data"))
		}
		numbers[i] = ai
	}

	fmt.Println("Enter k:")
	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		log.Fatal(errors.New("incorrect data"))
	}

	kMax, err := max.FindKMax(&numbers, k)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kMax)
}
