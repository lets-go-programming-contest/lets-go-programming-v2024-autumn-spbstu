package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	content := `package main

import "fmt"

func printGreeting() {
	fmt.Println("Hi, i'm a generated text")
}`

	outPath := filepath.Join("greeting.go")

	err := os.WriteFile(outPath, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Error generating 'greting.go': %w", err)
	}

	fmt.Println("Generating greeting.go")
}
