package main

import (
	"fmt"

	"github.com/kirill.romanchuk/task-3/strtructs"
)

func main() {
	config := strtructs.Config{}

	err := config.Parse()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(config.InputFile)
		fmt.Println(config.OutputFile)
	}
}
