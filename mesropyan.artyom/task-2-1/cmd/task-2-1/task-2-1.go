package main

import (
	"fmt"

	optimaltemp "github.com/artem6554/task-2-1/optimalTemp"
)

func main() {
	if err := optimaltemp.OptimalTemp(); err != nil {
		fmt.Println(err.Error())
	}
}
