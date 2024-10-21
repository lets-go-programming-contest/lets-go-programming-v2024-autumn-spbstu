package internal

import (
	"bufio"
	"container/heap"
	"io"
	"strconv"
	hp "task-2-2/internal/heap"
	"task-2-2/internal/reader"
)

func Application(in io.Reader, out io.Writer) error {
	n, err := reader.IntMainDataRead(in)
	if err != nil {
		return err
	}
	meals, err := reader.IntMealRead(in, n)
	if err != nil {
		return err
	}
	h := &hp.Heap{}
	heap.Init(h)
	for _, meal := range meals {
		heap.Push(h, meal)
	}
	k, err := reader.IntMainDataRead(in)
	if err != nil {
		return err
	}
	var mealK int
	for idx := 0; idx < k; idx++ {
		mealK = heap.Pop(h).(int)
	}
	output := bufio.NewWriter(out)
	defer output.Flush()
	str_value := strconv.Itoa(mealK)
	output.WriteString(str_value)
	output.WriteByte('\n')
	return nil
}
