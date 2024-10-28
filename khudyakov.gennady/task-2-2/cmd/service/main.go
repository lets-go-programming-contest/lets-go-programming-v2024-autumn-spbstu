package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KRYST4L614/task-2-2/internal/dishes"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	dishesData, dishNum, err := dishes.ReadData(reader)

	if err != nil {
		panic(err.Error())
	}

	result, err := dishes.FindRequestedDish(dishesData, dishNum)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}
