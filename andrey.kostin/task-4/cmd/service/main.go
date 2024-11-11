package main

import (
	"fmt"
	"log"

	reader "github.com/IDevFrye/task-4/internal/io"
)

var (
	firstTest = 1
	lastTest  = 2
)

func main() {
	choice, err := reader.GetNumber("Выберите режим работы (1 — потокобезопасный, 2 — небезопасный): ", firstTest, lastTest)
	if err != nil {
		log.Fatal(err)
	}
	if choice == 1 {

	} else if choice == 2 {

	} else {
		fmt.Println("Некорректный выбор. Введите число 1 или 2.")
	}
}
