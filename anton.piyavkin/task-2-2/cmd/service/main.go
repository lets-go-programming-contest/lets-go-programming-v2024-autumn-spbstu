package main

import (
	"fmt"
	"os"

	"github.com/Piyavva/task-2-2/internal/getDish"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
        os.Exit(1)
	}
	var dish int
    dishes := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&dish)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
            os.Exit(1)
		}
		dishes[i] = dish
	}
    var k int
	_, err = fmt.Scan(&k)
    if err != nil || k < 1 || k > n {
        fmt.Fprint(os.Stderr, "invalid k\n")
        os.Exit(1)
    }
    res := getDish.GetKDish(dishes, k)
	fmt.Println(res)
}