package main

func Func(a, b int) int {
	for i := range b {
		a -= i
	}
	return a
}

func main() {
	a, b := 10, 100
	Func(a, b)
}