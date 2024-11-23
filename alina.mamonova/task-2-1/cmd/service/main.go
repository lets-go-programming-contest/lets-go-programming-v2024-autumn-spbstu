package main

import (
	"bufio"
	"fmt"
	"github.com/hahapathetic/task-2-1/internal/departments_handler"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Optimal temperature handler!")
	departments_handler.Start(reader)
}
