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
	ErrInvalidArraySize = errors.New("размер массива - плохое значение")
	ErrInvalidK         = errors.New("k - плохое значение")
	ErrInvalidInput     = errors.New("ввод некорректен")
)

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

func ReadSliceSize() (int, bool, error) {
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

func ReadSliceElements(sliceSize int) ([]int, bool, error) {
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

func ReadKthElement(sliceSize int) (int, bool, error) {
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
