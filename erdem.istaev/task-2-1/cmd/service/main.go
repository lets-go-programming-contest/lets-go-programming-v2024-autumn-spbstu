package service

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrIncorrectNumber = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minHeap := make(MinHeap, k)
	for i := 0; i < k; i++ {
		minHeap[i] = nums[i]
	}

	heap.Init(&minHeap)
	for i := k; i < len(nums); i++ {
		if minHeap[0] < nums[i] {
			minHeap[0] = nums[i]
			heap.Fix(&minHeap, 0)
		}
	}
	return minHeap[0]
}

func main() {
	var n, k int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество элементов в массиве: ")
	n, err := readQuantity(reader)
	if err != nil || n <= 0 {
		fmt.Println(err)
		return
	}

	nums := make([]int, n)
	fmt.Print("Введите элементы в массиве: ")
	for i := 0; i < n; i++ {
		nums[i], err = readQuantity(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Print("Введите k-ый наибольший элемент: ")
	k, err = readQuantity(reader)
	if err != nil || k <= 0 || len(nums) < k {
		fmt.Println(err)
		return
	}

	fmt.Printf("Результат: %d", findKthLargest(nums, k))

}

func readQuantity(reader *bufio.Reader) (int, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	remSpc := strings.TrimSpace(str)
	var res int
	if res, err = strconv.Atoi(remSpc); err != nil {
		return 0, ErrIncorrectNumber
	}
	return res, nil
}
