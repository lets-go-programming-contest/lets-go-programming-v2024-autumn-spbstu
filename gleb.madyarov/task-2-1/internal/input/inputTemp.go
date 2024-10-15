package input

import (
	"fmt"
	"log"
)

const ERR_INT = "Error, expected integer"
const ERR_SIGN = "Error sign"

func AddNumber() int {
	var numb int
	_, err := fmt.Scanln(&numb)
	if err != nil {
		log.Fatal(ERR_INT)
	}
	return numb
}

func AddTempSign() string {
	var tempSign string
	fmt.Scan(&tempSign)
	if tempSign != "<=" && tempSign != ">=" {
		log.Fatal(ERR_SIGN)
	}
	return tempSign
}
