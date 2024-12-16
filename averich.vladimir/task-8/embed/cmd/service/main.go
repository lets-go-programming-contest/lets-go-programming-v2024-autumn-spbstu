package main

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/*.hash
var folder embed.FS

func main() {
	fmt.Println("File as string:", fileString)
	fmt.Println("File as byte slice:", string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println("Content of file1.hash:", string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println("Content of file2.hash:", string(content2))
}
