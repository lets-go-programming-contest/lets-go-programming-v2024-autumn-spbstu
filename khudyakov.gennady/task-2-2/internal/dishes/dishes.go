package dishes

import (
	"bufio"
	"container/heap"
	"strconv"
	"strings"

	"github.com/KRYST4L614/task-2-2/internal/intheap"
	"github.com/KRYST4L614/task-2-2/internal/io"
)

func FindRequestedDish(dishesData []int, requestedDish int) (int, error) {

	if requestedDish > len(dishesData) {
		return 0, RequestedDishError{
			DishNum:    requestedDish,
			DishesSize: len(dishesData),
		}
	} else if requestedDish < 0 {
		return 0, NegativeDishNumberError{DishNum: requestedDish}
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

	return dishes, k, nil
}
