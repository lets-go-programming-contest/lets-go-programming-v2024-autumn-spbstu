package main

import (
	"fmt"

	"github.com/mrqiz/task-4/internal/parking"
)

func main() {
	park := parking.NewParking(5)
	fmt.Printf("okay so there are %d slots in the parking\n", park.Capacity())
}
