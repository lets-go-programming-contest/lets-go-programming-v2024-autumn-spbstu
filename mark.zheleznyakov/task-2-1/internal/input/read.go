package input

import (
	"fmt"
	"strconv"
)

func ReadInt(label string) int {
  var input string

	for {
		fmt.Printf("now reading %s: ", label)
		fmt.Scanln(&input)

		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("this is not a number")
			continue
		}

    return n 
	}
}

