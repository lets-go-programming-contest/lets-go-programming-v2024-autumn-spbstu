package main

import (
	"embed"
)

//go:embed folder/file.txt
var fileStr string

//go:embed folder/file.txt
var fileBytes []byte

//go:embed folder/*.hash
var folder embed.FS

func main() {
	println(fileStr)
	println(string(fileBytes))
	content1, _ := folder.ReadFile("folder/file1.hash")
	println(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	println(string(content2))
}
