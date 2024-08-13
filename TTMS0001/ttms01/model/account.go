package model

type Account struct {
	Username   string
	tickets    []*Ticket
	TotalCount int
}
