package dao

import (
	"strconv"
	"ttms01/model"
	"ttms01/utils"
)

func GetMovieSessionByMovieName(moviename string) (*model.Page, error) {
	sql := "select * from movie_session where movie = ?"
	row, _ := utils.Db.Query(sql, moviename)
	var moviesessions []*model.MovieSession
	for row.Next() {
		moviesession := &model.MovieSession{}
		_ = row.Scan(&moviesession.ShowCinema, &moviesession.ShowScreen, &moviesession.ShowTime, &moviesession.ShowMovie, &moviesession.ShowInfo, &moviesession.Price)
		moviesession.State = moviesession.JudgeState()
		moviesession.Remaining = moviesession.Count()
		//fmt.Println("GetMovieSessionByMovieName", moviesession.State)
		//fmt.Println("handlerzhongde", moviesession)
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

func ModifyShowSessionSeat(moviesession *model.MovieSession, newinfo string) error {
	sql := "update movie_session set info=? where cinema=? and screenroom=? and movietime=? and movie=? and price=?"
	_, err := utils.Db.Exec(sql, newinfo, moviesession.ShowCinema, moviesession.ShowScreen, moviesession.ShowTime, moviesession.ShowMovie, moviesession.Price)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMovieSession(moviesession *model.MovieSession) error {
	sql := "delete from movie_session where cinema = ? and screenroom = ? and movietime = ? and movie = ?"
	utils.Db.Exec(sql, moviesession.ShowCinema, moviesession.ShowScreen, moviesession.ShowTime, moviesession.ShowMovie)
	return nil
}

func SaveMovieSession(moviesession *model.MovieSession) error {
	sql := "insert into movie_session(cinema,screenroom,movietime,movie,info,price) values (?,?,?,?,?,?)"
	utils.Db.Exec(sql, moviesession.ShowCinema, moviesession.ShowScreen, moviesession.ShowTime, moviesession.ShowMovie, moviesession.ShowInfo, moviesession.Price)
	return nil
}

func GetMoviesessionByCinemaAndScreen(cinema string, screenroom string) ([]*model.MovieSession, error) {
	sql := "select * from movie_session where cinema=? and screenroom=? "
	rows, _ := utils.Db.Query(sql, cinema, screenroom)
	var moviesessions []*model.MovieSession
	for rows.Next() {
		moviesession := &model.MovieSession{}
		rows.Scan(&moviesession.ShowCinema, &moviesession.ShowScreen, &moviesession.ShowTime, &moviesession.ShowMovie, &moviesession.ShowInfo, &moviesession.Price)
		//fmt.Println("GetMoviesessionByCinemaAndScreen", moviesession)
		moviesessions = append(moviesessions, moviesession)
	}
	return moviesessions, nil
}
