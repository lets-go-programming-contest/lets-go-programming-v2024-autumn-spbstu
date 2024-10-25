package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("You need to use the flag -config to input yaml file")
	}
	if os.Args[1] != "-config" {
		fmt.Println("You need to use the flag -config to input yaml file")
		os.Exit(1)
	}
	yamlFile, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
