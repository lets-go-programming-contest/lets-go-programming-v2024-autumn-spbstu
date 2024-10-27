package main

import (
	"fmt"
	"log"

	"github.com/Mmmakskl/task-3/internal/config"
	"github.com/Mmmakskl/task-3/internal/logic"
)

func main() {

	configPath, err := config.ParseFlag()
	if err != nil {
		log.Fatal(err)
	}

	conf := config.Config{}
	if err = config.LoadConfig(configPath, &conf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running with config %v\n", conf)

	valutes := logic.ValCurs{}
	if err := logic.Parser(conf.InputFile, &valutes); err != nil {
		log.Fatal(err)
	}

	if err := logic.WriteJSON(conf.OutputFile, &valutes); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nData successfully written to JSON")
}
