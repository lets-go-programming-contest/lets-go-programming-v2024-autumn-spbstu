package tests

import (
	"log"
	safeBook "main/internal/safeBooking"
	"sync"
)

func RunSafeTest(gorNumb int, ticketNumb int) {
	log.Println("------Start safe test-------")
	var wg sync.WaitGroup
	tbs := safeBook.NewTicketBookingSystem(ticketNumb)

	for i := 0; i < gorNumb; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tbs.BookTicket() {
				log.Println("tickets succesfully booked!")
			} else {
				log.Println("ticketd are sold out.")
			}
		}()
	}

	wg.Wait()

	log.Printf("tickets left: %d\n", tbs.GetAvailableTickets())
	log.Println("------end safe test-------")
}
