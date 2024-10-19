package internal

import (
	"fmt"
	"log"
	"os"
)

func ComputeTemp(sign string, temp *int, maxTemp *int, minTemp *int) {
	if sign != ">=" && sign != "<=" {
		log.Fatal("Invalid temperature")
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
}
