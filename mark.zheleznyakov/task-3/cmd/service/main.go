package main

import (
	"fmt"
	"log"

	"github.com/mrqiz/task-3/internal/config"
)

func main() {
	cLocation := config.ReadConfigFlag()
	c := config.ConfigFile{}

	err := config.Parse(&c, cLocation)
	if err != nil {
		log.Panicf("err: %v", err)
	}
	
	fmt.Println(c.InputFile, c.OutputFile)
}
