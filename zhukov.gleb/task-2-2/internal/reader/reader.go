package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidArraySize = errors.New("Array size - bad value")
	ErrInvalidK         = errors.New("k - bad value")
	ErrInvalidInput     = errors.New("Incorrect input")
)

type ConsoleReader struct{}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (c *ConsoleReader) ReadSliceSize() (int, bool, error) {
	var sliceSize int
	fmt.Println("Введите размер массива:")
	fmt.Print("> ")
	exit, err := readData(&sliceSize)
	if exit || err != nil {
		return 0, exit, err
	}
	if sliceSize <= 0 {
		return 0, false, ErrInvalidArraySize
	}
	return sliceSize, false, nil
}

func (c *ConsoleReader) ReadSliceElements(sliceSize int) ([]int, bool, error) {
	var slice []int
	fmt.Println("Введите элементы массива через пробел:")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, false, err
	}
	input = strings.TrimSpace(input)

	elements := strings.Split(input, " ")
	if len(elements) != sliceSize {
		return nil, false, ErrInvalidArraySize
	}

	for _, elem := range elements {
		number, err := strconv.Atoi(elem)
		if err != nil {
			return nil, false, ErrInvalidInput
		}
		slice = append(slice, number)
	}

	return slice, false, nil
}

func (c *ConsoleReader) ReadKthElement(sliceSize int) (int, bool, error) {
	var k int
	fmt.Println("Введите k-ый больший элемент ( >0 ):")
	fmt.Print("> ")
	exit, err := readData(&k)
	if exit || err != nil {
		return 0, exit, err
	}
	if k-1 >= sliceSize || k <= 0 {
		return 0, false, ErrInvalidK
	}
	return k, false, nil
}

func readData(number *int) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	input = strings.TrimSpace(input)
	if input == "выход" {
		return true, nil
	}
	*number, err = strconv.Atoi(input)
	if err != nil {
		return false, ErrInvalidInput
	}
	return false, nil
}
