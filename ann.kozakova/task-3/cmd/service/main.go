package main

import (
	"github.com/nutochk/task-3/internal/encoding"
	"github.com/nutochk/task-3/internal/parsing"
)

func main() {
	config := parsing.ConfigParsing()
	valCurs := parsing.InputParsing(config)
	encoding.JsonEncoding(config, valCurs)
}
