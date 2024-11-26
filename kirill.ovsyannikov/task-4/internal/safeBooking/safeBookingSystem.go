package safeBooking

import (
	"sync"
)

type TicketBookingSystem struct {
	tickets int
	mu      sync.Mutex
}

func NewTicketBookingSystem(totalTickets int) *TicketBookingSystem {
	return &TicketBookingSystem{
		tickets: totalTickets,
	}
}

func (tbs *TicketBookingSystem) BookTicket() bool {
	tbs.mu.Lock()
	defer tbs.mu.Unlock()

	if tbs.tickets > 0 {
		tbs.tickets--
		return true
	}
	return false
}

func (tbs *TicketBookingSystem) GetAvailableTickets() int {
	tbs.mu.Lock()
	defer tbs.mu.Unlock()
	return tbs.tickets
}
