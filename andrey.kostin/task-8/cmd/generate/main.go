//go:generate mockgen -source=main.go -destination=./mocks/mock_processor.go -package=mocks
package main

import "fmt"

type StringProcessor interface {
	Process(input string) string
}

func main() {
	fmt.Println("Пример использования go:generate")
}
