package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		dep        int
		emp        int
		tempSign   string
		tempValue  int
		tempLower  []int
		tempHigher []int
	)
	fmt.Scanln(&dep)
	for i := 1; i <= dep; i++ {
		fmt.Scanln(&emp)
		for j := 1; j <= emp; j++ {
			fmt.Scan(&tempSign)
			fmt.Scanln(&tempValue)
			if tempSign == "<=" {
				tempLower = append(tempLower, tempValue)
			}
			if tempSign == ">=" {
				tempHigher = append(tempHigher, tempValue)
			}
			sort.Ints(tempLower)
			sort.Ints(tempHigher)
			if len(tempLower) == 0 {
				fmt.Println(tempHigher[len(tempHigher)-1])
				continue
			}
			if len(tempHigher) == 0 {
				fmt.Println(tempLower[0])
				continue
			}
			if tempLower[0] >= tempHigher[len(tempHigher)-1] {
				fmt.Println(tempHigher[len(tempHigher)-1])
			} else {
				fmt.Println(-1)
			}
		}
		tempLower = nil
		tempHigher = nil
	}
}
