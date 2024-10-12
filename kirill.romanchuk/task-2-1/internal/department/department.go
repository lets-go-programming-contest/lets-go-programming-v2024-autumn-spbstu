package department

import (
	"fmt"
	"math"

	"github.com/kirill.romanchuk/task-2-1/internal/utils"
)

type Department struct {
	lowerBound int
	upperBound int
}

func NewDepartment() Department {
	return Department{lowerBound: utils.LowerBound, upperBound: utils.UpperBound}
}

func (d *Department) ManageTemperature() {
	k := utils.ReadIntNum("Введите количество сотрудников (1-2000): ", 1, 2000)

	for j := 0; j < k; j++ {
		condition, temperature := utils.ReadConditionAndTemperature()

		switch condition {
		case ">=":
			d.lowerBound = int(math.Max(float64(d.lowerBound), float64(temperature)))
		case "<=":
			d.upperBound = int(math.Min(float64(d.upperBound), float64(temperature)))
		}

		if d.lowerBound > d.upperBound {
			fmt.Println("Невозможно подобрать температуру для этого отдела\n", -1)
			break
		} else {
			fmt.Printf("Подходящая температура для отдела: %d\n", d.lowerBound)
		}
	}
}
