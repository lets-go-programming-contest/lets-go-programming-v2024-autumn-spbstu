package main

import (
	"fmt"
	"os"

	"github.com/kirill.romanchuk/task-2-1/internal/department"
	"github.com/kirill.romanchuk/task-2-1/internal/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
			os.Exit(1)
		}
	}()

	n := utils.ReadIntNum("Введите количество отделов (1-2000): ", 1, 2000)
	departments := make([]department.Department, n)

	for i := range departments {
		departments[i] = department.NewDepartment()
	}

	for i := range departments {
		departments[i].ManageTemperature()
	}
}
