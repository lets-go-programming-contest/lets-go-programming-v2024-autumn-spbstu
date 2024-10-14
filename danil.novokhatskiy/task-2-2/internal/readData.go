package internal

import (
	"fmt"
	"log"
)

func ReadData() ([]int, int, error) {
	fmt.Print("Enter the number of dishes: ")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatal(err)
	}
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		var tmp int
		_, err := fmt.Scan(&tmp)
		if err != nil || tmp > 10000 || tmp < -10000 {
			fmt.Println("Invalid input")
		}
		arr[i] = tmp
	}
	fmt.Println("Enter k value: ")
	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		log.Fatal(err)
	}
	return arr, k, err
}
