package main

import (
	_ "embed"
	"fmt"
)

//go:embed wings.png
var image []byte

//go:embed wings.png
var image2 string

func main() {
	fmt.Println(image)
	if string(image) == image2 {
		println("true")
	}
}
