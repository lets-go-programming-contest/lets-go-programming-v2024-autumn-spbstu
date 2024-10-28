package main

import (
	"fmt"
	"task-2-2/internal/findDish"
	"task-2-2/internal/input"
)

func main() {

	dishNum, dishes, err := input.InputData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(findDish.FindDish(dishNum, *dishes))

}
