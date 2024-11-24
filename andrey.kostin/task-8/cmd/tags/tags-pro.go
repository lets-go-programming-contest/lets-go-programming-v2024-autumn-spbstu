//go:build pro
// +build pro

package main

import (
	"fmt"
	"strings"
)

type ProProcessor struct{}

func (p ProProcessor) Process(input string) {
	words := strings.Fields(input)
	fmt.Println("Pro версия активирована.")
	fmt.Println("Количество слов:", len(words))

	countChars := len([]rune(input))
	fmt.Println("Количество символов:", countChars)
}

func init() {
	processor = ProProcessor{}
}
