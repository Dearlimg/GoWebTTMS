package model

type Page struct {
	Movies          []*Movie
	Movie           *Movie
	Cinema          []*Cinema
	ShowSession     []*MovieSession
	ShowSession1    *MovieSession
	Session         *Session
	Tickets         []Ticket
	ComingMovies    []*Movie
	ClassicMovies   []*Movie
	BoxOfficeMovies []*Movie
	ExpectedMovies  []*Movie
	TopMovies       []*Movie
	Comments        []*Comment
	Introductions   *Introduction

	PageNo      int64
	PageSize    int64
	TotalPageNo int64
	TotalRecord int64
	IsLogin     bool
	IsAdmin     bool
	Username    string
	Showinfo    []int
	Seatinfo    []int
	SumPrice    float64
	Message     string
}

func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
