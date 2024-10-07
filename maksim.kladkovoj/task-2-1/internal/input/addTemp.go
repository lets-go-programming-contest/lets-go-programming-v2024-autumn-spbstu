package input

import (
	"fmt"
	"log"

	"github.com/Mmmakskl/task-2-1/pkg/errors"
)

func AddNumber() int {
	var n int

	fmt.Print("Enter the number of departments: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal(errors.ErrInput)
	}

	return n
}
