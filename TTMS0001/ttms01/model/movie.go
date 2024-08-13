package model

type Movie struct {
	MovieName string
	ActorName string
	Time      string
	Score     float64
	BoxOffice int64
	Genre     string
	Area      string
	Age       string
	ImgPath   string
	Duration  string
	Cinema    string
	Showtime  string
}

type Introduction struct {
	MovieName string
	Intro     string
}
