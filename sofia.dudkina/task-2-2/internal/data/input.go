package data

import (
	"errors"
	"fmt"
)

func InputInt(x *int) error {
	_, err := fmt.Scan(x)
	if err != nil {
		return errors.New("incorrect input")
	} else {
		return nil
	}
}

func InputSlice(slice []int) error {
	var temp int
	for i := 0; i < len(slice); i++ {
		err := InputInt(&temp)
		if err != nil {
			return err
		}
		slice[i] = temp
	}
	return nil
}

func GetTaskCondition() ([]int, int, error) {
	var N, k int
	err := InputInt(&N)
	if err != nil {
		return nil, 0, err
	}
	dishes := make([]int, N)
	err = InputSlice(dishes)
	if err != nil {
		return nil, 0, err
	}
	err = InputInt(&k)
	if err != nil {
		return nil, 0, err
	}
	return dishes, k, nil
}
