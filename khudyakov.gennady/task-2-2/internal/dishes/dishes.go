package dishes

import (
	"bufio"
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/KRYST4L614/task-2-2/internal/intheap"
	"github.com/KRYST4L614/task-2-2/internal/io"
)

func requestedDishErr(dishNum, dishesSize int) error {
	return fmt.Errorf("порядковый номер <%v> k-го по предпочтению блюда больше, чем общее кол-во блюд <%v>", dishNum, dishesSize)
}

func negativeDishErr(dishNum int) error {
	return fmt.Errorf("порядковый номер <%v> k-го по предпочтению блюда должен быть больше нуля", dishNum)
}

func FindRequestedDish(dishesData []int, requestedDish int) (int, error) {

	if requestedDish > len(dishesData) {
		return 0, requestedDishErr(requestedDish, len(dishesData))
	} else if requestedDish < 0 {
		return 0, negativeDishErr(requestedDish)
	}

	intHeap := &intheap.IntHeap{}
	heap.Init(intHeap)

	for _, dish := range dishesData {
		heap.Push(intHeap, dish)
	}

	for range requestedDish - 1 {
		heap.Pop(intHeap)
	}

	return heap.Pop(intHeap).(int), nil
}

func ReadData(reader *bufio.Reader) ([]int, int, error) {
	_, err := io.ReadInt(reader)
	if err != nil {
		return nil, 0, err
	}

	inputString, err := reader.ReadString('\n')
	if err != nil {
		return nil, 0, err
	}

	dishesString := strings.Split(strings.TrimSpace(inputString), " ")

	dishes := make([]int, len(dishesString))

	for i := range dishesString {
		dish, err := strconv.ParseInt(dishesString[i], 10, 0)
		if err != nil {
			return nil, 0, err
		}
		dishes[i] = int(dish)
	}

	k, err := io.ReadInt(reader)
	if err != nil {
		return nil, 0, err
	}

	return dishes, int(k), nil
}
