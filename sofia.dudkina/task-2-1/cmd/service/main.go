package main

import "fmt"

func main() {
	var N, K int
	var minT, maxT, T int
	var op string
	fmt.Scan(&N)
	for ; N > 0; N-- {
		fmt.Scan(&K)
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
