package main

import (
	"embed"
)

//go:embed text.txt
var fileString string

//go:embed file1.hash
var fileHash embed.FS

func main() {
	println(fileString)
	content1, _ := fileHash.ReadFile("file1.hash")
	println(string(content1))
}
