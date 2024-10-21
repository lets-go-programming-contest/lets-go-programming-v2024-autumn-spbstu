package data

import (
	"github.com/sssidkn/task-2-2/pkg/maxk"
)

func Solution() (int, error) {
	dishes, k, err := GetTaskCondition()
	if err != nil {
		return 0, err
	}
	res, err := maxk.Find(dishes, k)
	if err != nil {
		return 0, err
	}
	return res, nil
}
