package model

import (
	"ttms01/utils"
)

type MovieSession struct {
	MovieSessionId string
	MovieId        string
	CinemaId       string
	ShowCinema     string
	ShowScreen     string
	ShowTime       string
	ShowMovie      string
	ShowInfo       string
	Price          float64
	State          bool
	Remaining      int
	MovieImgPath   string
	Duration       string
}

func (session *MovieSession) JudgeState() bool {
	for i := 0; i < len(session.ShowInfo); i++ {
		if session.ShowInfo[i] == '0' {
			return false
		}
	}
	return true
}

func (session *MovieSession) Count() int {
	res := 0
	for i := 0; i < len(session.ShowInfo); i++ {
		if session.ShowInfo[i] == '0' {
			res++
		}
	}
	return res
}

func (session *MovieSession) GetMovieDuration(ShowMovie string) string {
	sql := "select Duration from movie where moviename=?"
	row := utils.Db.QueryRow(sql, ShowMovie)
	res := ""
	row.Scan(&res)
	//fmt.Println("GetMovieTime", res)
	return res
}
