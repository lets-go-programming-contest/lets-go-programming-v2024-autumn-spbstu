package main

import "fmt"

func main() {

	var num1 float32
	fmt.Print("Введите первое число: ")
	_, err1 := fmt.Scan(&num1)
	if err1 != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	var op string
	fmt.Print("Выберите операцию (+, -, *, /): ")
	_, errOp := fmt.Scan(&op)
	if errOp != nil || (op != "+" && op != "-" && op != "/" && op != "*") {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}

	var num2 float32
	fmt.Print("Введите второе число: ")
	_, err2 := fmt.Scan(&num2)
	if err2 != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	switch op {
	case "+":
		fmt.Printf("Результат %v + %v = %v", num1, num2, num1+num2)
	case "*":
		fmt.Printf("Результат %v * %v = %v", num1, num2, num1*num2)
	case "-":
		fmt.Printf("Результат %v - %v = %v", num1, num2, num1-num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно")
		} else {
			fmt.Printf("Результат %v/%v=%v", num1, num2, num1/num2)
		}
	}
}
