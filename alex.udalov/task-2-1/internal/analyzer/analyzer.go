package analyzer

import (
	"fmt"
	"log"
)

const (
	minDeptCount    = 1
	maxDeptCount    = 1000
	minWorkersCount = 1
	maxWorkersCount = 1000
	minTemp         = 15
	maxTemp         = 30
)

func AnalyzeDeptCount() int {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil || num < minDeptCount || num > maxDeptCount {
		log.Fatal("Ошибка некорректное количество отделов")
	}

	return num
}

func AnalyzeWorkersCount() int {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < minWorkersCount || n > maxWorkersCount {
		log.Fatal("Ошибка некорректное количество сотрудников")
	}

	return n
}

func AnalyzeTemp() (string, int) {
	var k int
	var oper string
	_, err := fmt.Scanf("%s %d\n", &oper, &k)
	if err != nil || k > maxTemp || k < minTemp || (oper != ">=" && oper != "<=") {
		log.Fatal("Ошибка: некорректный ввод данных")
	}
	return oper, k
}
