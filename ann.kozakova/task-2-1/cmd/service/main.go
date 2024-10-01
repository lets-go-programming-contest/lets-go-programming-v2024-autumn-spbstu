package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter numbers separated by spaces:")
	scanner.Scan()
	line := scanner.Text()
	nums := strings.Split(line, " ")
	var numbers []int
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(errors.New("incorrect data"))
		}
		numbers = append(numbers, num)
	}
	fmt.Println(numbers)

	fmt.Println("Enter k:")
	var k int
	_, err := fmt.Scan(&k)
	if err != nil {
		log.Fatal(errors.New("incorrect data"))
	}

	kMax, err := findKMax(&numbers, k)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kMax)
}

func minValue(nums *[]int) (int, int, error) {
	if len((*nums)) == 0 {
		return 0, 0, errors.New("Cannot detect a minimum value in an empty slice")
	}

	min := (*nums)[0]
	index := 0
	for i, v := range *nums {
		if v < min {
			min = v
			index = i
		}
	}

	return min, index, nil
}

func findKMax(nums *[]int, k int) (int, error) {
	if k <= 1 || k > len((*nums)) {
		return 0, errors.New("incorrect k")
	}

	var maxNums []int

	for _, num := range *nums {
		if len(maxNums) < k {
			maxNums = append(maxNums, num)
		}

		min, index, _ := minValue(&maxNums)
		if num > min {
			maxNums = append(maxNums[0:index], maxNums[index+1:]...)
			maxNums = append(maxNums, num)
		}
	}
	min, _, _ := minValue(&maxNums)
	return min, nil
}
