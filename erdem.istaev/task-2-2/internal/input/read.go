package input

import (
	"errors"
	"fmt"
)

var (
	ErrIncorrectNumber = errors.New("Incorrect number")
)

func ReadData() (int, []int, int, error) {
	var n, ai, k int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		return 0, nil, 0, ErrIncorrectNumber
	}

	dishes := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&ai)
		if err != nil {
			return 0, nil, 0, ErrIncorrectNumber
		}
		dishes[i] = ai
	}

	_, err = fmt.Scan(&k)
	if err != nil || k <= 0 || k > n {
		return 0, nil, 0, ErrIncorrectNumber
	}

	return k, dishes, n, nil
}
