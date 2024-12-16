package main

func isOperator(operator string) bool {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	return operators[operator]
}

func main() {
	var operation string = "/"
	isOperator(operation)
}

// go tool compile -W main.go
// go tool compile -S main.go
// go tool compile -o main.o main.go
