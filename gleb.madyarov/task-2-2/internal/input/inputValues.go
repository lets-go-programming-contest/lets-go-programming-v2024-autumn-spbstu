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
	//fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		ai = CheckNumb()
		//fmt.Scan(&ai)
		numb = append(numb, ai)
	}
	k = CheckNumb()
	//fmt.Scanln(&k)
	return numb, k
}
