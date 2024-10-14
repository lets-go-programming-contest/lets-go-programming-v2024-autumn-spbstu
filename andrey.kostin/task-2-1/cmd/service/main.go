package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp = 15
	maxTemp = 30
	minVal  = 1
	maxVal  = 1000
)

func GetInt(prompt string, minVal int, maxVal int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()

		if num, err := strconv.Atoi(input); err == nil {
			if num <= maxVal && num >= minVal {
				return num
			} else {
				fmt.Println(errors.ErrorIncorectIntBounds)
			}
		} else {
			fmt.Println(errors.ErrorIncorectInt)
		}
	}
}

func GetTempCondition(iter int, prompt string) (string, int) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(iter, ". ", prompt)
		scanner.Scan()
		input := scanner.Text()

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println(errors.ErrorIncorectCondSpace)
			continue
		}

		cond := parts[0]
		if cond != "<=" && cond != ">=" {
			fmt.Println(errors.ErrorIncorectCondComp)
			continue
		}

		tempStr := parts[1]
		temp, err := strconv.Atoi(tempStr)
		if err != nil {
			fmt.Println(errors.ErrorIncorectCondTemp)
			continue
		}

		return cond, temp
	}
}

func main() {
	countOfDepts := GetInt("Введите количество отделов: ", minVal, maxVal)

	for i := 0; i < countOfDepts; i++ {
		countOfEmps := GetInt("Введите количество сотрудников в отделе: ", minVal, maxVal)
		for j := 0; j < countOfEmps; j++ {
			cond, temp := GetTempCondition(j+1, "Введите комфортную температуру (например, '>= 30'): ")
		}
	}
}
