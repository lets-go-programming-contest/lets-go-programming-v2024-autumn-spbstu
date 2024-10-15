package calculate

import (
	"fmt"

	"github.com/Mmmakskl/task-2-1/internal/input"
)

func OptimalTemp(k int) error {
	var (
		topTemp  int = 30
		downTemp int = 15
	)

	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		operator, temperature, err := input.AddTemperature()
		if err != nil {
			return err
		}

		switch operator {
		case input.LessOrEqual:
			if temperature <= topTemp {
				topTemp = temperature
			}
		case input.Less:
			if temperature-1 < topTemp {
				topTemp = temperature - 1
			}
		case input.GreaterOrEqual:
			if temperature >= downTemp {
				downTemp = temperature
			}
		case input.Greater:
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
	return nil
}
