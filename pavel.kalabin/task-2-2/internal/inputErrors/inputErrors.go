package inputErrors

import "fmt"

type NRangeError struct {}

func (e NRangeError) Error() string {
    return fmt.Sprintf("N must be from 1 to 10000")
}

type AIRangeError struct {}

func (e AIRangeError) Error() string {
    return fmt.Sprintf("Dish priority must be from -10000 to 10000")
}

type KRangeError struct {
    N int
}

func (e KRangeError) Error() string {
    return fmt.Sprintf("k value must be from 1 to %d", e.N) 
}

