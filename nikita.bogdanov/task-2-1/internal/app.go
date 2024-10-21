package internal

import (
	"io"
	reader "task-2-1/internal/reader"
	writer "task-2-1/internal/writer"
)

func Application(in io.Reader, out io.Writer) error {
	n, err := reader.Read_main_data(in)
	if err != nil {
		return err
	}
	for department := 0; department < n; department++ {
		lowBoundTemp := 15
		upBoundTemp := 30
		k, err := reader.Read_main_data(in)
		if err != nil {
			return err
		}
		for employee := 0; employee < k; employee++ {
			number, comp, err := reader.Read_conditioner_data(in)
			if err != nil {
				return err
			}
			if comp == ">=" && number >= lowBoundTemp {
				lowBoundTemp = number
			} else if comp == "<=" && number <= upBoundTemp {
				upBoundTemp = number
			}
			if upBoundTemp >= lowBoundTemp {
				writer.Write_optional_temperature(out, lowBoundTemp)
			} else {
				writer.Write_optional_temperature(out, -1)
			}
		}
	}
	return nil
}
