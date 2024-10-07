package calculate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/Mmmakskl/task-2-1/pkg/errors"
)

const expectedMatchCount = 3

func OptimalTemp(int) {
	var (
		k        int
		topTemp  int = 100
		downTemp int = -100
		operator string
		value    string
	)

	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of employees: ")
	_, err := fmt.Scanln(&k)
	if err != nil {
		log.Fatal(errors.ErrInput)
	}

	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		value, err = in.ReadString('\n')
		if err != nil {
			log.Fatal(errors.ErrInput)
		}

		re, err := regexp.Compile(`([<>]=?)\s*(\d+)`)
		if err != nil {
			log.Fatal(errors.ErrRegexp)
		}

		matches := re.FindStringSubmatch(value)
		if len(matches) < expectedMatchCount {
			log.Fatal(errors.ErrInput)
		}

		operator = matches[1]
		temperature, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(errors.ErrTemp)
		}

		switch operator {
		case "<=":
			if temperature <= topTemp {
				topTemp = temperature
			}
		case "<":
			if temperature-1 < topTemp {
				topTemp = temperature - 1
			}
		case ">=":
			if temperature >= downTemp {
				downTemp = temperature
			}
		case ">":
			if temperature+1 > downTemp {
				downTemp = temperature + 1
			}
		}

		if topTemp < downTemp {
			fmt.Println(-1)
			break
		} else if downTemp == -100 && i == 0 {
			fmt.Println(topTemp)
		} else {
			fmt.Println(downTemp)
		}
	}
}
