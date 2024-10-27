package main

import (
	"fmt"
	"log"
	"os"

	"github.com/katagiriwhy/task-3/internal"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("You need to use the flag -config to input yaml file")
	}
	if os.Args[1] != "-config" {
		fmt.Println("You need to use the flag -config to input yaml file")
		os.Exit(1)
	}
	yaml := internal.ReadFlag()

	cfg := internal.Config{}
	curr := internal.Currencies{}

	err := internal.ReadAndParseConfig(&cfg, yaml)
	if err != nil {
		log.Fatal(err)
	}
	err = internal.ParseXml(&curr, cfg.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	err = internal.Convert(&curr, cfg.OutputDir)
}
