package main

import (
	_ "embed"
	"fmt"
	"reflect"
)

//go:embed pic.jpg
var fileByte []byte

//go:embed pic.jpg
var fileString string

func main() {
	fmt.Println(reflect.DeepEqual(fileString, string(fileByte)))
}
