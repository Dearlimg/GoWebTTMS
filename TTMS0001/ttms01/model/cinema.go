package model

type Cinema struct {
	CinemaId      int
	CinemaName    string
	Place         string
	Rank          string
	ScreenRoomNum int
	State         bool
}

type CinemaSearch struct {
	District []string
	Rank     []string
}
