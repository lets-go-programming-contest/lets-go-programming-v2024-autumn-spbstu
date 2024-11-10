package internal

import (
	"errors"
	"io"
)

func GetTemp(k int, out io.Writer, maxTemp *int, minTemp *int) error {
	for i := 0; i < k; i++ {
		temp, sign, err := ReadData()
		if err != nil {
			return err
		}
		if sign != ">=" && sign != "<=" {
			return errors.New("Invalid sign")
		}
		if sign == "<=" && temp < *maxTemp {
			*maxTemp = temp
		}
		if sign == ">=" && temp > *minTemp {
			*minTemp = temp
		}
		if *maxTemp >= *minTemp {
			WriteInt(*minTemp, out)
		} else {
			return ErrorTemp{}
		}
	}
	return nil
}
