package client

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/kirill.romanchuk/task-4/structures/performance"
)

type Client struct {
	ID int
}

func (c *Client) ReserveSeat(perf *performance.Performance, mutex *sync.Mutex, wGroup *sync.WaitGroup) {
	defer wGroup.Done()
	resultChan := make(chan bool)
	row, col := rand.Intn(perf.Rows), rand.Intn(perf.Cols)
	go perf.ReserveSeat(row, col, mutex, resultChan)
	select {
	case result := <-resultChan:
		if result {
			fmt.Printf("Клиент ID %d: Место [%d, %d] успешно забронировано.\n", c.ID, row, col)
		} else {
			fmt.Printf("Клиент ID %d: Место [%d, %d] уже занято.\n", c.ID, row, col)
		}
	case <-time.After(20 * time.Second):
		fmt.Printf("Клиент ID %d: Время ожидания истекло при попытке забронировать место [%d, %d].\n", c.ID, row, col)
	}
}
