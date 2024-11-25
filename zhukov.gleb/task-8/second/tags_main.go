package main

import "fmt"

var sliceElem = []string{
	"first",
	"second",
}

func main() {
	for _, e := range sliceElem {
		fmt.Println(e)
	}
}
