package internal

import "fmt"

func ReadData() (int, string, error) {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return 0, "", err
	}
	var temp int
	_, err = fmt.Scan(&temp)
	if err != nil {
		return 0, "", err
	}
	return temp, str, nil
}
