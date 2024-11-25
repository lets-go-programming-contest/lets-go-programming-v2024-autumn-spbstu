package main

import (
	"fmt"
	"sync"

	"erdem.istaev/task-4/internal/library"
)

func main() {
	bookRequests := []string{"Harry Potter", "Green Mile", "Lord of the rings", "The Shawshank Redemption",
		"The Master and Margarita", "Green Mile"}
	lib := library.NewLibrary([]string{"Harry Potter", "Green Mile", "Sherlock Holmes"})

	var wg sync.WaitGroup
	for _, request := range bookRequests {
		wg.Add(1)
		go func(book string) {
			defer wg.Done()
			lib.BorrowBook(book)
		}(request)
	}
	wg.Wait()

	fmt.Println("Available books after borrowing:", lib.AvailableBooks())

	lib.ReturnBook("Harry Potter")
	lib.ReturnBook("Green Mile")

	fmt.Println("Available books after returning:", lib.AvailableBooks())
}
