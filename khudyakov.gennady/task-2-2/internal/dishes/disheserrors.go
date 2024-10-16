package dishes

import "fmt"

func requestedDishErr(dishNum, dishesSize int) error {
	return fmt.Errorf("порядковый номер <%v> k-го по предпочтению блюда больше, чем общее кол-во блюд <%v>", dishNum, dishesSize)
}

func negativeDishErr(dishNum int) error {
	return fmt.Errorf("порядковый номер <%v> k-го по предпочтению блюда должен быть больше нуля", dishNum)
}
