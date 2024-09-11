package model

type Movie struct {
	MovieId      string
	MovieName    string
	ActorName    string
	Showtime     string
	Score        float64
	BoxOffice    int64
	Genre        string
	Area         string
	Age          string
	ImgPath      string
	Duration     string
	Introduction string
	State        string
}

type Introduction struct {
	MovieName string
	Intro     string
}

type MovieSearch struct {
	Genre  []string
	Region []string
	Decade []string
}
