package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KRYST4L614/task-2-1/internal/io"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, err := io.ReadInt(reader)
	if err != nil {
		fmt.Println(err.Error())
	}

	max := int64(30)
	min := int64(15)

	for range n {
		currentMax := max
		currentMin := min

		k, err := io.ReadInt(reader)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for range k {
			term, temperature, err := io.ReadTerm(reader)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if temperature < 15 || temperature > 30 {
				fmt.Printf("Некорректное значение температуры. Допускаются значения от %v до %v\n", min, max)
				return
			}

			switch term {
			case ">=":
				currentMin = temperature
			case "<=":
				currentMax = temperature
			default:
				fmt.Println("Некорректное условие для температуры. Допустимы: <= или >=")
				return
			}
			if currentMin > currentMax {
				fmt.Println(-1)
				return
			}
			fmt.Println(currentMin)
		}
	}
}
