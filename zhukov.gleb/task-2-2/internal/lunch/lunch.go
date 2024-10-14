package lunch

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"container/heap"
	myHeap "task-2-2/internal/heap"
)

type ConsoleLunch struct{}

func NewConsoleLunch() *ConsoleLunch {
	return &ConsoleLunch{}
}

func readData(number *int) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	input = strings.TrimSpace(input)
	if input == "выход" {
		return true, nil
	}
	*number, err = strconv.Atoi(input)
	return false, err
}

func readSliceSize() (int, bool, error) {
	var sliceSize int
	fmt.Println("Введите размер массива:")
	fmt.Print("> ")
	exit, err := readData(&sliceSize)
	if exit || err != nil {
		return 0, exit, err
	}
	if sliceSize <= 0 {
		return 0, false, fmt.Errorf("размер массива - плохое значение")
	}
	return sliceSize, false, nil
}

func readSliceElements(sliceSize int) ([]int, bool, error) {
	slice := make([]int, sliceSize)
	fmt.Println("Вводите числа массива через новую строку:")
	for i := 0; i < sliceSize; i++ {
		fmt.Print("> ")
		exit, err := readData(&slice[i])
		if exit {
			return nil, true, nil
		}
		if err != nil {
			return nil, false, err
		}
	}
	return slice, false, nil
}

func readKthElement(sliceSize int) (int, bool, error) {
	var k int
	fmt.Println("Введите k-ый больший элемент ( >0 ):")
	fmt.Print("> ")
	exit, err := readData(&k)
	if exit || err != nil {
		return 0, exit, err
	}
	if k-1 >= sliceSize || k <= 0 {
		return 0, false, fmt.Errorf("k - плохое значение")
	}
	return k, false, nil
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

func (c *ConsoleLunch) Run() {
	fmt.Println("Поиск k-го наибольшего числа массива")
	fmt.Println("'выход' - для завершения")

	for {
		sliceSize, exit, err := readSliceSize()
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		slice, exit, err := readSliceElements(sliceSize)
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		k, exit, err := readKthElement(sliceSize)
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		fmt.Printf("Результат: %d\n", findKthElement(slice, k))
	}
}
