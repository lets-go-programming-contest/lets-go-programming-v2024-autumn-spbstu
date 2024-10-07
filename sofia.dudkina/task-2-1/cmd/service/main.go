package main

import (
	"fmt"
	"os"
)

func main() {
	var N, K int
	var minT, maxT, T int
	var op string
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for ; N > 0; N-- {
		_, err = fmt.Scan(&K)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		minT = 15
		maxT = 30
		for ; K > 0; K-- {
			fmt.Scan(&op, &T)
			switch op {
			case ">=":
				if T >= 15 {
					minT = max(T, minT)
				}
			case "<=":
				maxT = min(T, maxT)
			}
			if minT > maxT {
				fmt.Println(-1)
			} else {
				println(minT)
			}
		}
	}
}
