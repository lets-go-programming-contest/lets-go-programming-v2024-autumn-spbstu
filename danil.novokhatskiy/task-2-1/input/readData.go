package input

import (
	"bufio"
	"fmt"
	"github.com/katagiriwhy/task-2-1/errors"
	"log"
	"os"
)

func inputNum() int {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal(errors.ErrInput)
	}
	return num
}

func inputTemp() (string, int) {
	var sign string
	in := bufio.NewReader(os.Stdin)
	sign, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(errors.ErrInput)
	}
	if !(sign == ">=" || sign == "<=") {
		log.Fatal(errors.ErrInput)
	}
	num := inputNum()
	return sign, num
}
