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

func ReadCondition() (string, int) {
	var sign string
	var temp int

	for {
		fmt.Printf("reading condition: ")
		
		_, err := fmt.Scan(&sign, &temp)

		if err != nil {
			fmt.Println("unable to read condition")
			continue
		}

		if sign != "<=" && sign != ">=" {
			fmt.Println("bad sign")
			continue
		}

		break
	}

	return sign, temp
}

