package input

import (
	"fmt"
	"strconv"
	"os"
)

func ReadInt() int {
  var input string

	fmt.Scanln(&input)

	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("this is not a number")
		os.Exit(1)
	}

	return n 
}

func ReadCondition() (string, int) {
	var sign string
	var temp int

	_, err := fmt.Scan(&sign, &temp)

	if err != nil {
		fmt.Println("unable to read condition")
		os.Exit(1)
	}

	if sign != "<=" && sign != ">=" {
		fmt.Println("bad sign")
		os.Exit(1)
	}

	return sign, temp
}

