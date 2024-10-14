package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	errproc "github.com/IDevFrye/task-2-1/internal/errors"
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
				fmt.Println(errproc.ErrorIncorectIntBounds)
			}
		} else {
			fmt.Println(errproc.ErrorIncorectInt)
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
			fmt.Println(errproc.ErrorIncorectCondSpace)
			continue
		}

		cond := parts[0]
		if cond != "<=" && cond != ">=" {
			fmt.Println(errproc.ErrorIncorectCondComp)
			continue
		}

		tempStr := parts[1]
		temp, err := strconv.Atoi(tempStr)
		if err != nil {
			fmt.Println(errproc.ErrorIncorectCondTemp)
			continue
		}

		return cond, temp
	}
}
