package main

import (
	"fmt"
	"os"

	"github.com/kirill.romanchuk/task-2-2/internal/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	n, values, k := utils.ReadInput()
	fmt.Print(n, values, k)
}
