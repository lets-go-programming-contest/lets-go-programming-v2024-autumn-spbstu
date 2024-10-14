package main

import (
	"github.com/IDevFrye/task-2-1/internal/bounds"
	internal "github.com/IDevFrye/task-2-1/internal/io"
)

const (
	minTemp = 15
	maxTemp = 30
	minVal  = 1
	maxVal  = 1000
)

func main() {
	countOfDepts := internal.GetInt("Введите количество отделов: ", minVal, maxVal)

	tempBounds := make([]bounds.TempBounds, countOfDepts)
	for i := 0; i < countOfDepts; i++ {
		tempBounds[i] = bounds.TempBounds{LowerBound: minTemp, UpperBound: maxTemp}
	}

	for i := 0; i < countOfDepts; i++ {
		countOfEmps := internal.GetInt("Введите количество сотрудников в отделе: ", minVal, maxVal)
		for j := 0; j < countOfEmps; j++ {
			cond, temp := internal.GetTempCondition(j+1, "Введите комфортную температуру (например, '>= 30'): ")
			tempBounds[i].EditBounds(cond, temp)
		}
	}
}
