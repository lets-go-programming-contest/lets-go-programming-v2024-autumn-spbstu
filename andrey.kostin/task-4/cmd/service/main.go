package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/IDevFrye/task-4/internal/io"
	"github.com/IDevFrye/task-4/internal/test"
)

var (
	firstTest = 1
	lastTest  = 2
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	choice, err := io.GetNumber("Выберите режим работы (1 — потокобезопасный, 2 — небезопасный): ", firstTest, lastTest, reader)
	if err != nil {
		log.Fatal(err)
	}
	if choice == 1 {
		test.SafeTest()
	} else if choice == 2 {
		test.UnsafeTest()
	} else {
		fmt.Println("Некорректный выбор. Введите число 1 или 2.")
	}
}
