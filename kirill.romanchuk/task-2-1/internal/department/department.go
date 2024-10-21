package department

import (
	"fmt"
	"math"

	"github.com/kirill.romanchuk/task-2-1/internal/utils"
)

const (
	LowerBound = 15
	UpperBound = 30
)

type Department struct {
	lowerBound int
	upperBound int
}

func NewDepartment() Department {
	return Department{lowerBound: LowerBound, upperBound: UpperBound}
}

func (d *Department) UpdateBounds(condition string, temperature int) error {
	switch condition {
	case ">=":
		if temperature > UpperBound {
			return fmt.Errorf("ошибка: температура не может быть больше %d", UpperBound)
		}
		d.lowerBound = int(math.Max(float64(d.lowerBound), float64(temperature)))
		return nil

	case "<=":
		if temperature < LowerBound {
			return fmt.Errorf("ошибка: температура не может быть меньше %d", LowerBound)
		}
		d.upperBound = int(math.Min(float64(d.upperBound), float64(temperature)))
		return nil

	default:
		panic(fmt.Errorf("ошибка: неверное условие '%s'", condition))
	}
}

func (d *Department) ManageTemperature() {
	k := utils.ReadIntNum("Введите количество сотрудников (1-2000): ", 1, 2000)

	for j := 0; j < k; j++ {
		condition, temperature := utils.ReadConditionAndTemperature()

		err := d.UpdateBounds(condition, temperature)

		if err != nil {
			fmt.Println(-1)
			break
		}

		if d.lowerBound > d.upperBound {
			fmt.Println(-1)
			break
		} else {
			fmt.Printf("Подходящая температура для отдела: %d\n", d.lowerBound)
		}
	}
}
