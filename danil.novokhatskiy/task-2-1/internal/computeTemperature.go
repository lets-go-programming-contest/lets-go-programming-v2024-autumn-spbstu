package internal

import (
	"errors"
	"fmt"
	"os"
)

func ComputeTemp(sign string, temp *int, maxTemp *int, minTemp *int) (err error) {
	if sign != ">=" && sign != "<=" {
		return errors.New("Invalid sign")
	}
	if sign == "<=" && *temp < *maxTemp {
		*maxTemp = *temp
	}
	if sign == ">=" && *temp > *minTemp {
		*minTemp = *temp
	}
	if *maxTemp >= *minTemp {
		fmt.Println(*minTemp)
	} else {
		fmt.Println(-1)
		os.Exit(1)
	}
	return nil
}
