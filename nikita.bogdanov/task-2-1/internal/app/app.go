package app

import (
	"io"

	reader "task-2-1/internal/reader"
	writer "task-2-1/internal/writer"
)

const OtherWiseVar = -1

func Application(in io.Reader, out io.Writer) error {
	n, err := reader.ReadMainData(in)
	if err != nil {
		return err
	}
	for department := 0; department < n; department++ {
		lowBoundTemp := reader.MinTemperature
		upBoundTemp := reader.MaxTemperature
		k, err := reader.ReadMainData(in)
		if err != nil {
			return err
		}
		for employee := 0; employee < k; employee++ {
			number, comp, err := reader.ReadConditionerData(in)
			if err != nil {
				return err
			}
			if comp == reader.Greater && number >= lowBoundTemp {
				lowBoundTemp = number
			} else if comp == reader.Less && number <= upBoundTemp {
				upBoundTemp = number
			}
			if upBoundTemp >= lowBoundTemp {
				writer.Write_optional_temperature(out, lowBoundTemp)
			} else {
				writer.Write_optional_temperature(out, OtherWiseVar)
			}
		}
	}
	return nil
}
