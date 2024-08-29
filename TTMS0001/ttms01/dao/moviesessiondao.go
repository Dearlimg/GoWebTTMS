package dao

import (
	"strconv"
	"ttms01/model"
	"ttms01/utils"
)

func GetMovieSessionByMovieName(moviename string) (*model.Page, error) {
	sql0 := "select movieid from movie where moviename = ?"
	row0 := utils.Db.QueryRow(sql0, moviename)

	sql := "select * from moviesession where movieid = ?"
	row, _ := utils.Db.Query(sql, row0)
	var moviesessions []*model.MovieSession
	for row.Next() {
		moviesession := &model.MovieSession{}
		_ = row.Scan(&moviesession.MovieSessionId, &moviesession.MovieId, &moviesession.CinemaId, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
		moviesession.State = moviesession.JudgeState()
		moviesession.Remaining = moviesession.Count()
		moviesessions = append(moviesessions, moviesession)
	}
	page := &model.Page{
		ShowSession: moviesessions,
	}
	return page, nil
}

func ParseInfo(date string) ([]int, error) {
	nums := make([]int, 0, 40)
	for i := 0; i < len(date); i++ {
		num, _ := strconv.Atoi(string(date[i]))
		nums = append(nums, num)
	}
	return nums, nil
}

//func ModifyShowSessionSeatById(moviesession *model.MovieSession, newinfo string) error {
//	sql := "update moviesession set showinfo=?"
//	_, err := utils.Db.Exec(sql, newinfo, moviesession.ShowCinema, moviesession.ShowScreen, moviesession.ShowTime, moviesession.ShowMovie, moviesession.Price)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func DeleteMovieSession(moviesession *model.MovieSession) error {
	sql := "update moviesession set state=0 where sessionid = ?"
	utils.Db.Exec(sql, moviesession.MovieSessionId)
	return nil
}

func SaveMovieSession(moviesession *model.MovieSession) error {
	sql := "insert into moviesession(movieid,cinemaid,showtime,showinfo,state,screenroom,price) values (?,?,?,?,?,?,?)"
	utils.Db.Exec(sql, moviesession.MovieId, moviesession.CinemaId, moviesession.ShowTime, moviesession.ShowInfo, moviesession.State, moviesession.ShowScreen, moviesession.Price)

	return nil
}

func GetMoviesessionByCinemaAndScreen(cinema string, screenroom string) ([]*model.MovieSession, error) {
	sql := "select * from movie_session where cinema=? and screenroom=? "
	rows, _ := utils.Db.Query(sql, cinema, screenroom)
	var moviesessions []*model.MovieSession
	for rows.Next() {
		moviesession := &model.MovieSession{}
		rows.Scan(&moviesession.MovieSessionId, &moviesession.MovieId, &moviesession.CinemaId, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
		moviesessions = append(moviesessions, moviesession)
	}
	return moviesessions, nil
}
