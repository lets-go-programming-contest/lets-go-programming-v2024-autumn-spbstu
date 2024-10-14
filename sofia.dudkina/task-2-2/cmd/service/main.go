package main

import (
	"fmt"
	"github.com/sssidkn/task-2-2/internal/data"
)

func main() {
	answer, err := data.Solution()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
}
