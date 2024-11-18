package main

import (
	tests "main/internal/tests"
)

var numbOfGors int = 150
var numbOfTickets int = 100

func main() {
	tests.RunSafeTest(numbOfGors, numbOfTickets)

	tests.RunUnSafeTest(numbOfGors, numbOfTickets)

}
