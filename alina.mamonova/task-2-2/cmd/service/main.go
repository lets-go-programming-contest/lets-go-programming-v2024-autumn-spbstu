package main

import (
	"bufio"
	"fmt"
	"github.com/hahapathetic/task-2-2/internal/heap-handler"
	"github.com/hahapathetic/task-2-2/internal/input"
	"github.com/hahapathetic/task-2-2/internal/table-heap"
	"os"
)

const (
	minInputBound = 1
	maxInputBound = 10000
)

func main() {
	fmt.Println("Welcome to Dishes analyzer!")
	reader := bufio.NewReader(os.Stdin)

	dishNum := input.ProcessNumInput("Enter the number of dishes: ", minInputBound, maxInputBound, reader)
	table := &table_heap.TableHeap{}

	heap_handler.FillHeap(dishNum, table, reader)
	k := input.ProcessNumInput("Enter the k: ", 1, dishNum, reader)
	result := heap_handler.ProcessHeap(dishNum, table, k)

	fmt.Println("Result dish:", result)
}
