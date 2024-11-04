package dishes

import "fmt"

type RequestedDishError struct {
	DishNum    int
	DishesSize int
}

func (err RequestedDishError) Error() string {
	return fmt.Sprintf("порядковый номер <%v> k-го по предпочтению блюда больше, чем общее кол-во блюд <%v>", err.DishNum, err.DishesSize)
}

type NegativeDishNumberError struct {
	DishNum int
}

func (err NegativeDishNumberError) Error() string {
	return fmt.Sprintf("порядковый номер <%v> k-го по предпочтению блюда должен быть больше нуля", err.DishNum)
}
