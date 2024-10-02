package main

import (
	"fmt"
	"math"
)

func getKel(arr []int, k int) int {
	var tmp []int = []int{}
	mn := math.MaxInt
	for i := 0; i < k; i++ {
		mn = min(mn, arr[i])
		tmp = append(tmp, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if mn < arr[i] {
			var temp int = arr[i]
			for j := 0; j < len(tmp); j++ {
				if mn == tmp[j] {
					tmp[j] = arr[i]
					mn = arr[i]
				} else {
					temp = min(temp, tmp[j])
				}
			}
			mn = temp
		}
	}
	return mn
}

func main() {
    var arr []int = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(getKel(arr, 4))
}