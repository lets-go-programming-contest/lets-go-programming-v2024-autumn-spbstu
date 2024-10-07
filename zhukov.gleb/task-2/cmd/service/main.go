package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadData(number *int) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	input = strings.TrimSpace(input)
	if input == "выход" {
		return true, nil
	}
	*number, err = strconv.Atoi(input)
	return false, err
}

func main() {
	fmt.Println("Поиск k-го наибольшего числа массива")
	fmt.Println("'выход' - для завершения")

LOOP:
	for {
		var sliceSize int
		var exit bool
		fmt.Println("Введите размер массива:")
		fmt.Print("> ")
		exit, err := ReadData(&sliceSize)
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if sliceSize <= 0 {
			fmt.Printf("размер массива - плохое значение\n")
			continue
		}

		slice := make([]int, sliceSize)
		fmt.Println("Вводите числа массива через новую строку:")
		for i := 0; i < sliceSize; i++ {
			fmt.Print("> ")
			exit, err := ReadData(&slice[i])
			if exit {
				return
			}
			if err != nil {
				fmt.Printf("%v\n", err)
				continue LOOP
			}
		}

		var k int
		fmt.Println("Введите k-ый больший элемент ( >0 ):")
		fmt.Print("> ")
		exit, err = ReadData(&k)
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if k-1 >= len(slice) || k <= 0 {
			fmt.Printf("k - плохое значение\n")
			continue
		}

		sort.Slice(slice, func(i, j int) bool {
			return slice[i] > slice[j]
		})

		fmt.Printf("Результат: %d\n", slice[k-1])
	}
}
