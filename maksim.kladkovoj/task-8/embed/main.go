package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed data/text.txt
var textFile string

//go:embed data/data.json
var jsonData []byte

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Print("Text data:\n")
	fmt.Println(textFile)

	var users []User
	err := json.Unmarshal(jsonData, &users)
	if err != nil {
		log.Fatalf("Failed parsin json: %w", err)
	}

	fmt.Println("\nUsers: ")
	for _, user := range users {
		fmt.Printf("%s is %d years old\n", user.Name, user.Age)
	}
}
