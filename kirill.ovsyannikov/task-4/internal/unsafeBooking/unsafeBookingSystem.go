package unsafeBooking

// TicketBookingSystem представляет собой систему бронирования билетов на концерт.
type TicketBookingSystem struct {
	tickets int
}

// NewTicketBookingSystem создает новую систему бронирования билетов.
func NewTicketBookingSystem(totalTickets int) *TicketBookingSystem {
	return &TicketBookingSystem{
		tickets: totalTickets,
	}
}

// BookTicket бронирует билет, если они еще доступны.
func (tbs *TicketBookingSystem) BookTicket() bool {
	if tbs.tickets > 0 {
		tbs.tickets--
		return true
	}
	return false
}

// GetAvailableTickets возвращает количество оставшихся билетов.
func (tbs *TicketBookingSystem) GetAvailableTickets() int {
	return tbs.tickets
}
