package control

import (
	"errors"
	"fmt"
	"strings"
	myReader "task-2-1/internal/reader"
)

var (
	ErrInvalidTempString = errors.New("bad temp request")
)

type ConsoleControl struct {
	reader myReader.ConsoleReader
}

func NewConsoleControl(reader myReader.ConsoleReader) *ConsoleControl {
	return &ConsoleControl{reader: reader}
}

func (c *ConsoleControl) Run() error {
	fmt.Println("Поиск оптимальной температуры для отдела")
	fmt.Println("'выход' - для завершения")

	fmt.Println("Введите количество отделов:")
	fmt.Print("> ")
	n, exit, err := c.reader.ReadNK()
	if exit {
		return nil
	}
	if err != nil {
		return fmt.Errorf("ошибка при чтении количества отделов: %v", err)
	}

	for i := 0; i < n; i++ {
		fmt.Println("Введите количество сотрудников:")
		fmt.Print("> ")
		k, exit, err := c.reader.ReadNK()
		if exit {
			return nil
		}
		if err != nil {
			return fmt.Errorf("ошибка при чтении количество сотрудников: %v", err)
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
			fmt.Print("> ")
			condition, exit, err := c.reader.ReadCondition()
			if exit {
				return nil
			}
			if err != nil {
				return fmt.Errorf("ошибка при чтении температуры: %v", err)
			}

			minTemp, maxTemp, err = updateTemperatureBounds(minTemp, maxTemp, condition, c)
			if err != nil {
				return fmt.Errorf("ошибка при работе с температурой: %v", err)
			}

			if minTemp <= maxTemp {
				fmt.Printf("%d\n", minTemp)
			} else {
				fmt.Println("-1")
			}
		}
	}
	return nil
}

func updateTemperatureBounds(minTemp, maxTemp int, condition string, c *ConsoleControl) (int, int, error) {
	parts := strings.Split(condition, " ")
	if len(parts) != 2 {
		return minTemp, maxTemp, ErrInvalidTempString
	}

	operator := parts[0]
	value, err := c.reader.ParseTemperature(parts[1])
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
