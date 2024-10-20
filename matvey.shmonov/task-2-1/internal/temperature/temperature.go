package temperature

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	MaxTemperature = 30
	MinTemperature = 15
)
const pattern = `^([<>]=)\s*([+-]?\d+)$`

func ReadInt() int {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return input
}

func ReadPrompt(b *bufio.Reader) string {
	prompt, _ := b.ReadString('\n')
	return strings.TrimSpace(prompt)
}

func ProcessInput(prompt string, minT, maxT *int) bool {
	re := regexp.MustCompile(pattern)
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
