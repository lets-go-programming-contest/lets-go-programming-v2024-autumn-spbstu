package calculate

import (
	"fmt"

	"github.com/Mmmakskl/task-2-1/internal/input"
)

func OptimalTemp(int) {
	var (
		topTemp  int = 100
		downTemp int = -100
	)

	fmt.Print("Enter the number of employees: ")
	k := input.AddNumber()

	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		operator, temperature := input.AddTemperature()

		switch operator {
		case "<=":
			if temperature <= topTemp {
				topTemp = temperature
			}
		case "<":
			if temperature-1 < topTemp {
				topTemp = temperature - 1
			}
		case ">=":
			if temperature >= downTemp {
				downTemp = temperature
			}
		case ">":
			if temperature+1 > downTemp {
				downTemp = temperature + 1
			}
		}

		if topTemp < downTemp {
			fmt.Println(-1)
			break
		} else if downTemp == -100 {
			fmt.Println(topTemp)
		} else {
			fmt.Println(downTemp)
		}
	}
}
