package main

import (
	"embed"
	"fmt"
)

//go:embed byte.txt
var fileByte string

//go:embed test.txt
var fileString string

//go:embed folder
var folder embed.FS

func main() {
	fmt.Println(fileByte)

	fmt.Println(fileString)

	file1, _ := folder.ReadFile("folder/file1.txt")
	fmt.Println(string(file1))

	file2, _ := folder.ReadFile("folder/file2.txt")
	fmt.Println(string(file2))

	file3, _ := folder.ReadFile("folder/file3.txt")
	fmt.Println(string(file3))
}
