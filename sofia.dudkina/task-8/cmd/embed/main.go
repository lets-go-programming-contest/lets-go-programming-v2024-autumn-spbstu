package main

import (
	_ "embed"
	"fmt"
)

//go:embed src/text1.txt
var text string

//go:embed src/text1.txt
var byteText []byte

func main() {
	fmt.Println(text)
	fmt.Println(string(byteText))
	if text == string(byteText) {
		fmt.Println("OK")
	}
}
