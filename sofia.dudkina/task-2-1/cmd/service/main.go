package main

import (
	"fmt"
	"github.com/sssidkn/task-2-1/internal/optimalt"
)

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println(err)
	} else {
		err = optimalt.Find(N)
		if err != nil {
			fmt.Println(err)
		}
	}
}
