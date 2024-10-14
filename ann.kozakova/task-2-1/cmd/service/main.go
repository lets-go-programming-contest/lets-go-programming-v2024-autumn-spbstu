package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/nutochk/task-2-1/internal/temp"
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
		topT := 30
		bottomT := 15
		for j := 0; j < workers; j++ {
			var sign string
			fmt.Scan(&sign)
			if !(sign == "<=" || sign == ">=") {
				log.Fatal(errors.New("incorret data"))
			}
			var temperature int
			_, err := fmt.Scan(&temperature)
			if err != nil {
				log.Fatal(errors.New("incorret data"))
			}
			res, err := temp.FindTemp(&topT, &bottomT, &temperature, sign)
			if err != nil {
				fmt.Println(-1)
				return
			}
			fmt.Println(res)
		}
	}
}
