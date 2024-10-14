package input

import (
	"fmt"
	"log"
)

func CheckNumb() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func InputVal() ([]int, int) {
	var (
		n    int
		ai   int
		k    int
		numb []int
	)
	n = CheckNumb()
	for i := 0; i < n; i++ {
		ai = CheckNumb()
		numb = append(numb, ai)
	}
	k = CheckNumb()
	return numb, k
}
