package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/Mmmakskl/task-2-1/pkg/error"
)

const expectedMatchCount = 3

func AddNumber() int {
	var number int

	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal(error.ErrInput)
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
		log.Fatal(error.ErrInput)
	}

	re, err := regexp.Compile(`([<>]=?)\s*(\d+)`)
	if err != nil {
		log.Fatal(error.ErrRegexp)
	}

	matches := re.FindStringSubmatch(value)
	if len(matches) < expectedMatchCount {
		log.Fatal(error.ErrInput)
	}

	operator = matches[1]
	temperature, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal(error.ErrTemp)
	}

	return operator, temperature
}
