package main

import (
	"fmt"
	"task-2-2/internal/io"
	"task-2-2/internal/rating"
)

func main() {

	if index, heapOfDishes, errOfInput := io.InputParameters(); errOfInput != nil {

		fmt.Errorf("Error: %w\n", &errOfInput)

	} else if result, errOfDetect := rating.DetectTheMostGuessedDish(index, heapOfDishes); errOfDetect != nil {

		fmt.Errorf("Error: %w\n", &errOfDetect)

	} else {

		fmt.Println(result)

	}

}
