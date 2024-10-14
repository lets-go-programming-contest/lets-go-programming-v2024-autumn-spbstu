package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const expectedMatchCount = 3

func AddNumber() int {
	var number int

	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal(ErrInput)
	}

	return number
}

func AddTemperature() (string, int) {
	var (
		operator string
	)
	in := bufio.NewReader(os.Stdin)

	value, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(ErrInput)
	}

	re, err := regexp.Compile(`([<>]=?)\s*(\d+)`)
	if err != nil {
		log.Fatal(ErrRegexp)
	}

	matches := re.FindStringSubmatch(value)
	if len(matches) < expectedMatchCount {
		log.Fatal(ErrInput)
	}

	operator = matches[1]
	temperature, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal(ErrTemp)
	}

	return operator, temperature
}
