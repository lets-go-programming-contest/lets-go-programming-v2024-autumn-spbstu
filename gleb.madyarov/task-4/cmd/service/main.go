package main

import (
	"fmt"

	"github.com/Madyarov-Gleb/task-4/internal/test"
)

func main() {
	var input int
	fmt.Print("Choose safe or unsafe test (1 or 2): ")
	fmt.Scanln(&input)
	if input == 1 {
		test.Test()
	} else if input == 2 {
		test.TestUnsafe()
	} else {
		fmt.Println("Wrong choice")
	}
}
