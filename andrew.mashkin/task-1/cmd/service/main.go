package main

import (
	"fmt"
	"log"
)

func main() {
	var first int
	var second int
	var operation string
	_, err := fmt.Scan(&first, &second)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Scan(&operation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Calcualte(first, second, operation))
}
