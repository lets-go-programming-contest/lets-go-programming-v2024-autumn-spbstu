//go:build !pro
// +build !pro

package main

import (
	"fmt"
	"strings"
)

type BasicProcessor struct{}

func (p BasicProcessor) Process(input string) {
	words := strings.Fields(input)
	fmt.Println("Базовая версия активирована.")
	fmt.Println("Количество слов:", len(words))
	fmt.Println("Подсчёт символов недоступен в базовой версии.")
}

func init() {
	processor = BasicProcessor{}
}
