package temperature

import (
	"errors"
	"fmt"
	"github.com/katagiriwhy/task-2-1/input"
)

func proccessTemperature() {
	maxTemp := 30
	minTemp := 15
	fmt.Println("Input number of employees: ")
	numEmployees := 
	if sign != ">=" && sign != "<=" {
		return 0, errors.New("sign must be either '>=' or '<='")
	}
	if sign == "<=" && *temp < maxTemp {
		maxTemp = *temp
		return maxTemp, nil
	}
	if sign == ">=" && *temp > minTemp {
		minTemp = *temp
		return minTemp, nil
	}
	if maxTemp >= minTemp {
		return minTemp, nil
	} else {
		return -1
	}
}
