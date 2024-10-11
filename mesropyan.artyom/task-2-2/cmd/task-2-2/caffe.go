package main

import (
	"fmt"
	"sort"

	"github.com/artem6554/task-2-2/chooseDish"
	"github.com/artem6554/task-2-2/readDishes"
)

func main() {
	dishes := readDishes.ReadDishes()
	sort.Stable(&dishes)
	chosenDish := chooseDish.ChooseDish(dishes)
	fmt.Println(chosenDish)
}
