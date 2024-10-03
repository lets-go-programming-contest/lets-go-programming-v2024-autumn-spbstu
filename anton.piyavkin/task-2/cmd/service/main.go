package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func getKel(arr []int, k int) (int, error) {
    if (k > len(arr)) {
        return 0, errors.New("k больше размера массива")
    }
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
	return mn, nil
}

func main() {
    fmt.Println("Введите количество элементов массива:")
    var size int
    _, err := fmt.Scan(&size)
    if (err != nil) {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("Введите массив:")
    var arr []int
    for i := 0; i < size; i++ {
        var tmp int
        _, err = fmt.Scan(&tmp)
        if (err != nil) {
            fmt.Println(err)
            os.Exit(1)
        }
        arr = append(arr, tmp)
    }
    fmt.Println("Введите k:")
    var k int
    _, err = fmt.Scan(&k)
    if (err != nil) {
        fmt.Println(err)
        os.Exit(1)
    }
    res, err := getKel(arr, k)
    if (err != nil) {
        fmt.Println(err)
        os.Exit(2)
    }
    fmt.Println(res)
}