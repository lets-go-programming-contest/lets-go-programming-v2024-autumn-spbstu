package main

import (
	_ "embed"
	"fmt"
)

//go:embed pic.jpg
var fileByte []byte

//go:embed pic.jpg
var fileString string

func main() {
	fmt.Println(fileString == string(fileByte))
}
