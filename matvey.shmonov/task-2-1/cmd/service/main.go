package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MaxTemperature = 30
	MinTemperature = 15
)
const pattern = `^([<>]=)\s*([+-]?\d+)$`

func main() {
	b := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(pattern)

	N := readInt()
	for range N {
		minT, maxT := MinTemperature, MaxTemperature

		K := readInt()
		for range K {
			prompt := readPrompt(b)
			if ok := processInput(prompt, re, &minT, &maxT); ok {
				fmt.Println(minT)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func readInt() int {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return input
}

func readPrompt(b *bufio.Reader) string {
	prompt, _ := b.ReadString('\n')
	return strings.TrimSpace(prompt)
}

func processInput(prompt string, re *regexp.Regexp, minT, maxT *int) bool {
	matches := re.FindStringSubmatch(prompt)
	if len(matches) != 3 {
		log.Fatal("Error: Could not process input. Please check that the data you entered is correct.")
	}
	t, err := strconv.Atoi(matches[2])
	if err != nil { // owerflow or smth like that
		log.Fatalf("Error: Could not convert string '%s' to integer: %v", matches[2], err)
	}

	switch matches[1] {
	case ">=":
		if *minT < t {
			*minT = t
		}
	case "<=":
		if *maxT > t {
			*maxT = t
		}
	default: // unreachable
		log.Panicf("Unknown operator: '%s'", matches[1])
	}
	return MinTemperature <= *minT && *minT <= *maxT && *maxT <= MaxTemperature
}
