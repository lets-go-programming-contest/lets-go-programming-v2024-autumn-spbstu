package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter count of sections:")
	var sections int
	_, err := fmt.Scan(&sections)
	if err != nil {
		log.Fatal(errors.New("incorret data"))
	}

	for i := 0; i < sections; i++ {
		fmt.Println("Enter count of workers:")
		var workers int
		_, err := fmt.Scan(&workers)
		if err != nil {
			log.Fatal(errors.New("incorret data"))
		}
		fmt.Println("Enter the temperature:")
		topT := 100
		bottomT := -100
		for j := 0; j < workers; j++ {
			var sign string
			fmt.Scan(&sign)
			if !(sign == "<=" || sign == ">=") {
				log.Fatal(errors.New("incorret data"))
			}
			var temp int
			_, err := fmt.Scan(&temp)
			if err != nil {
				log.Fatal(errors.New("incorret data"))
			}
			if sign == "<=" && temp < topT {
				topT = temp
			}
			if sign == ">=" && temp > bottomT {
				bottomT = temp
			}
			if topT >= bottomT {
				fmt.Println(bottomT)
			} else {
				fmt.Println(-1)
				break
			}
		}
	}
}
