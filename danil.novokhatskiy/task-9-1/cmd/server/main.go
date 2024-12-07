package main

import (
	"fmt"
	"os"
	"task-9-1/internal/config"
)

func main() {
	configFile, err := os.Open("internal/config/config.yaml")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()
	cfg, err := config.Unmarshal(configFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", cfg)
}
