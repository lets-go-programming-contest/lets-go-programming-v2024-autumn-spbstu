package main

import _ "embed"

//go:embed src/text1.txt
var text string

//go:embed src/text2.txt
var byteText []byte

func main() {
	println(text)
	println(string(byteText))
}
