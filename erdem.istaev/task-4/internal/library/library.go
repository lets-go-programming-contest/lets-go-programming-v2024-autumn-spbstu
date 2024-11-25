package library

import (
	"fmt"
	"sync"
)

type Library struct {
	books []string
	mu    sync.Mutex
}

func NewLibrary(books []string) *Library {
	return &Library{books: books}
}

func (l *Library) BorrowBook(book string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i, b := range l.books {
		if b == book {
			fmt.Printf("Borrowed: %s\n", book)
			l.books = append(l.books[:i], l.books[i+1:]...)
			return
		}
	}

	fmt.Printf("Book not available: %s\n", book)
}

func (l *Library) ReturnBook(book string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.books = append(l.books, book)
	fmt.Printf("Returned: %s\n", book)
}

func (l *Library) AvailableBooks() []string {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.books
}
