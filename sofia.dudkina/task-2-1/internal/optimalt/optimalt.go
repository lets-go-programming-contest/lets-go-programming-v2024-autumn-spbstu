package optimalt

import (
	"errors"
	"fmt"
)

func Find(N int) error {
	var K int
	var minT, maxT, T int
	var op string
	for ; N > 0; N-- {
		_, err := fmt.Scan(&K)
		if err != nil {
			return errors.New("incorrect K")
		}
		minT = 15
		maxT = 30
		for ; K > 0; K-- {
			_, err = fmt.Scan(&op)
			if err != nil {
				return err
			}
			_, err = fmt.Scan(&T)
			if err != nil {
				return errors.New("incorrect T")
			}
			switch op {
			case ">=":
				if T >= 15 {
					minT = max(T, minT)
				}
			case "<=":
				maxT = min(T, maxT)
			default:
				return errors.New("invalid sign of comparison")
			}
			if minT > maxT {
				fmt.Println(-1)
			} else {
				fmt.Println(minT)
			}
		}
	}
	return nil
}
