package main

import (
	"sync"

	"github.com/kirill.romanchuk/task-4/structures/client"
	"github.com/kirill.romanchuk/task-4/structures/performance"
	"github.com/kirill.romanchuk/task-4/utils/readers"
)

func main() {
	var wGroup sync.WaitGroup
	var mutex sync.Mutex
	rows, cols := readers.ReadIntNum("Введите количисво рядов (1-3) ", 1, 3), readers.ReadIntNum("Введите количисво мест в ряду (1-3)", 1, 3)
	perm := performance.CreatePerformance(rows, cols)
	perm.DisplaySeats()

	numClients := readers.ReadIntNum("Введите количисво людей желающих забронировать место (1-16) ", 1, 16)
	wGroup.Add(numClients)
	for i := 1; i <= numClients; i++ {
		client := client.Client{ID: i}
		go client.ReserveSeat(&perm, &mutex, &wGroup)
	}

	wGroup.Wait()
	perm.DisplaySeats()
}
