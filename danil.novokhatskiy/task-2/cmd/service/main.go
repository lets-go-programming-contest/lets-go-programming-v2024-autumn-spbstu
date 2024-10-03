package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа, разделенные пробелами:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")

	var numbersArray []int
	for _, numberStr := range numbers {
		number, _ := strconv.Atoi(numberStr)
		numbersArray = append(numbersArray, number)
	}
	fmt.Println("Введите k: ")
	fmt.Print("> ")
	var k int
	_, err := fmt.Fscanln(reader, &k)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Считанные числа:", numbersArray)
}
