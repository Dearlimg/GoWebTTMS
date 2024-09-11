package model

type Ticket struct {
	TicketId       string
	UserId         string
	MovieSessionId string
	State          string
	Owner          string
	Movie          string
	Cinema         string
	Screen         string
	Seat           string
	Time           string
	Price          float64
}
