package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type TextProcessor interface {
	Process(input string)
}

var processor TextProcessor

func main() {
	if processor == nil {
		fmt.Println("Не задан процессор обработки текста!")
		return
	}

	fmt.Println("Введите текст:")

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatalf("Ошибка при вводе: %v\n", scanner.Err())
		return
	}
	input := scanner.Text()

	processor.Process(input)
}
