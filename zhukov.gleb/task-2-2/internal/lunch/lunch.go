package lunch

import (
	"container/heap"
	"fmt"
	myHeap "task-2-2/internal/heap"
	myReader "task-2-2/internal/reader"
)

type ConsoleLunch struct {
	reader myReader.ConsoleReader
}

func NewConsoleLunch(reader myReader.ConsoleReader) *ConsoleLunch {
	return &ConsoleLunch{reader: reader}
}

func findKthElement(slice []int, k int) int {
	data := &myHeap.IntHeap{}
	heap.Init(data)

	for _, e := range slice {
		heap.Push(data, e)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(data)
	}

	return heap.Pop(data).(int)
}

func (c *ConsoleLunch) Run() error {
	fmt.Println("Поиск k-го наибольшего числа массива")
	fmt.Println("'выход' - для завершения")

	for {
		sliceSize, exit, err := c.reader.ReadSliceSize()
		if exit {
			return nil
		}
		if err != nil {
			return fmt.Errorf("ошибка при чтении размера массива: %w", err)
		}

		slice, exit, err := c.reader.ReadSliceElements(sliceSize)
		if exit {
			return nil
		}
		if err != nil {
			return fmt.Errorf("ошибка при чтении элементов массива: %w", err)
		}

		k, exit, err := c.reader.ReadKthElement(sliceSize)
		if exit {
			return nil
		}
		if err != nil {
			return fmt.Errorf("ошибка при чтении индекса k: %w", err)
		}

		fmt.Printf("Результат: %d\n", findKthElement(slice, k))
	}
}
