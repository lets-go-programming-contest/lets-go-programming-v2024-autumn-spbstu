package tests

import (
	"log"
	unSafeBook "main/internal/unsafeBooking"
	"sync"
)

func RunUnSafeTest(gorNumb int, ticketNumb int) {
	log.Println("------Start unsafe test-------")
	var wg sync.WaitGroup
	tbs := unSafeBook.NewTicketBookingSystem(ticketNumb)

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
	log.Println("------end unsafe test-------")
}
