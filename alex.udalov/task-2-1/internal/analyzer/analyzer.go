package analyzer

import (
	"fmt"
	"log"
)

func AnalyzeDeptCount() int {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil || num < 1 || num > 1000 {
		log.Fatal("Ошибка некорректное количество отделов")
	}
	return num
}

func AnalyzeWorkersCount() int {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 || n > 1000 {
		log.Fatal("Ошибка некорректное количество сотрудников")
	}

	return n
}

func AnalyzeTemp() (string, int) {
	var k int
	var oper string
	_, err := fmt.Scanf("%s %d\n", &oper, &k)
	if err != nil || k > 30 || k < 15 || (oper != ">=" && oper != "<=") {
		log.Fatal("Ошибка: некорректный ввод данных")
	}
	return oper, k
}
