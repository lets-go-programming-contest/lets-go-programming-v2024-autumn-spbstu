package control

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"task-2-1/internal/reader"
)

var (
	ErrInvalidTempString = errors.New("плохой запрос температуры")
)

type ConsoleControl struct{}

func NewConsoleControl() *ConsoleControl {
	return &ConsoleControl{}
}

func (c *ConsoleControl) Run() {
	fmt.Println("Поиск оптимальной температуры для отдела")
	fmt.Println("'выход' - для завершения")

	fmt.Println("Введите количество отделов:")
	fmt.Print("> ")
	n, exit, err := reader.ReadNK()
	if exit {
		return
	}
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for i := 0; i < n; i++ {
		fmt.Println("Введите количество сотрудников:")
		fmt.Print("> ")
		k, exit, err := reader.ReadNK()
		if exit {
			return
		}
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
			fmt.Print("> ")
			condition, exit, err := reader.ReadCondition()
			if exit {
				return
			}
			if err != nil {
				log.Fatalf("%v\n", err)
			}

			minTemp, maxTemp, err = updateTemperatureBounds(minTemp, maxTemp, condition)
			if err != nil {
				fmt.Println("-1")
				continue
			}

			if minTemp <= maxTemp {
				fmt.Printf("%d\n", minTemp)
			} else {
				fmt.Println("-1")
			}
		}
	}
}

func updateTemperatureBounds(minTemp, maxTemp int, condition string) (int, int, error) {
	parts := strings.Split(condition, " ")
	if len(parts) != 2 {
		return minTemp, maxTemp, ErrInvalidTempString
	}

	operator := parts[0]
	value, err := reader.ParseTemperature(parts[1])
	if err != nil {
		return minTemp, maxTemp, err
	}

	switch operator {
	case ">=":
		if value > minTemp {
			minTemp = value
		}
	case "<=":
		if value < maxTemp {
			maxTemp = value
		}
	default:
		return minTemp, maxTemp, ErrInvalidTempString
	}

	return minTemp, maxTemp, nil
}
