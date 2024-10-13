package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"erdem.istaev/task-2-1/internal/errors"
	"erdem.istaev/task-2-1/internal/structure"
)

func ReadInt(reader *bufio.Reader) (int, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	remSpc := strings.TrimSpace(str)
	var res int
	if res, err = strconv.Atoi(remSpc); err != nil {
		return 0, errors.ErrIncorrectNumber
	}
	return res, nil
}

func ReadCondition(reader *bufio.Reader) (string, int, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, err
	}
	remSpc := strings.TrimSpace(str)
	parts := strings.Split(remSpc, " ")
	if len(parts) != 2 {
		return "", 0, errors.ErrIncorrectSeparator
	}

	if parts[0] != ">=" && parts[0] != "<=" {
		return "", 0, errors.ErrIncorrectComparsionOp
	}

	var temp int
	if temp, err = strconv.Atoi(parts[1]); err != nil {
		return "", 0, errors.ErrIncorrectNumber
	}

	return parts[0], temp, nil
}

func SetTemperature(reader *bufio.Reader, n int, bounds []structure.TemperatureBounds) {
	var condition string
	var temp int
	for i := 0; i < n; i++ {
		fmt.Print("Введите количество сотрудников: ")
		k, err := ReadInt(reader)
		if err != nil {
			fmt.Println(errors.ErrIncorrectNumber)
		}

		for j := 0; j < k; j++ {
			fmt.Print("Укажите температуру: ")
			condition, temp, err = ReadCondition(reader)
			if err != nil {
				fmt.Println(err)
				return
			}

			bounds[i].UpdateBounds(condition, temp)
			if bounds[i].IsValid() {
				fmt.Println(bounds[i].Lower)
			} else {
				fmt.Println(-1)
				return
			}
		}
	}
}
