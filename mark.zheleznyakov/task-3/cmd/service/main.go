package main

import (
	"fmt"
	"log"

	"github.com/mrqiz/task-3/internal/config"
)

var notRealConfigFileContents = `
input-file: "/tmp/money_in.yaml"
output-file: "/tmp/money_out.json"
`

func main() {
	c := config.ConfigFile{}
	err := config.ParseFromString(notRealConfigFileContents, &c)
	if err != nil {
		log.Fatalf("err: your config is cooked: %v", err)
	}
	fmt.Println(c.InputFile, c.OutputFile)
}
