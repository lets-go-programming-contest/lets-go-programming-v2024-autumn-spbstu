package main

import (
	"fmt"
	"github.com/hahapathetic/task-2-2/internal/heap-handler"
	"github.com/hahapathetic/task-2-2/internal/input"
	"github.com/hahapathetic/task-2-2/internal/table-heap"
)

const (
	minInputBound = 1
	maxInputBound = 10000
)

func main() {
	fmt.Println("Welcome to Dishes analyzer!")

	dishNum := input.ProcessNumInput("Enter the number of dishes: ", minInputBound, maxInputBound)
	table := &table_heap.TableHeap{}

	heap_handler.FillHeap(dishNum, table)
	result := heap_handler.ProcessHeap(dishNum, table)

	fmt.Println("Result dish:", result)
}
