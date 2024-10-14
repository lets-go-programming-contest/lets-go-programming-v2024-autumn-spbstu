package read

import (
	"fmt"
)

func ReadNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	for {
		if err != nil {
			fmt.Println("unable to read a number")
			continue
		}

		break
	}
	return n
}
