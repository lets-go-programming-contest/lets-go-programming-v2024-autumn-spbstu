package calculate

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var errInput error = errors.New("Input error")

func optimalTemp(int) {
	var (
		k        int
		topTemp  int = 100
		downTemp int = -100
		operator string
		value    string
	)

	in := bufio.NewReader(os.Stdin) //Создание переменной для чтения данных из потока ввода /в дальнейшем нужно для чтения строки/
	fmt.Print("Enter the number of employees: ")
	_, err := fmt.Scanln(&k)
	if err != nil {
		log.Fatal(errInput)
	}

	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		value, err = in.ReadString('\n')             //Чтение строки содержащей пробелы заканчивающейся символом конца строки '\n'
		re := regexp.MustCompile(`([<>]=?)\s*(\d+)`) //Регулярное выражение для проверки строки на нужный вид
		matches := re.FindStringSubmatch(value)      //Разделение value на совпадения [вся строка][оператор][число]
		if len(matches) < 3 {                        //Проверка на правильность ввода,в массиве должно быть 3 переменных
			log.Fatal(errInput)
		}

		operator = matches[1]
		temperature, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal("Error converting temperature")
		}

		switch operator { //Вычисление нужной температуры
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
		} else if downTemp == -100 && i == 0 { //Если переменная down не поменялась, то выводим topTemp
			fmt.Println(topTemp)
		} else {
			fmt.Println(downTemp)
		}
	}
}
