package main

var oper = []string{
	"*",
}

func main() {
	var res float64

	for _, o := range oper {
		res = operate(5, 2, o)
		println(res)
	}
}
