package dao

import (
	"fmt"
	"strconv"
	"ttms01/model"
	"ttms01/utils"
)

//func GetMovieSessionByMovieId(movieId string) (*model.Page, error) {
//	sql := `
//		SELECT
//			moviesession.moviesessionid,
//			movie.moviename,
//			cinema.cinemaname,
//			moviesession.showtime,
//			moviesession.showinfo,
//			moviesession.state,
//			moviesession.screenroom,
//			moviesession.price
//		FROM
//			moviesession
//		JOIN
//			movie ON movie.movieid = moviesession.movieid
//		JOIN
//			cinema ON cinema.cinemaid = moviesession.cinemaid
//		WHERE
//			moviesession.movieid = ?
//	`
//
//	fmt.Println(sql, movieId)
//
//	rows, err := utils.Db.Query(sql, movieId)
//	if err != nil {
//		return nil, fmt.Errorf("failed to execute query: %w", err)
//	}
//	defer rows.Close()
//
//	var moviesessions []*model.MovieSession
//	for rows.Next() {
//		moviesession := &model.MovieSession{}
//		err := rows.Scan(
//			&moviesession.MovieSessionId,
//			&moviesession.ShowMovie,
//			&moviesession.ShowCinema,
//			&moviesession.ShowTime,
//			&moviesession.ShowInfo,
//			&moviesession.State,
//			&moviesession.ShowScreen,
//			&moviesession.Price,
//		)
//		if err != nil {
//			return nil, fmt.Errorf("failed to scan row: %w", err)
//		}
//		moviesession.State = moviesession.JudgeState()
//		moviesession.Remaining = moviesession.Count()
//		fmt.Println("12312", moviesession)
//		moviesessions = append(moviesessions, moviesession)
//	}
//
//	page := &model.Page{
//		ShowSession: moviesessions,
//	}
//	return page, nil
//}

func GetMovieSessionByMovieId(movieId string) (*model.Page, error) {
	sql := "select movie.moviename,cinema.cinemaname,moviesession.showtime,moviesession.showinfo,moviesession.state,moviesession.screenroom,moviesession.price from moviesession join movie on movie.movieid=moviesession.movieid join cinema on cinema.cinemaid=moviesession.cinemaid where moviesession.movieid=?"
	row, _ := utils.Db.Query(sql, movieId)
	//fmt.Println(" GetMovieSessionByMovieId1", sql, movieId)
	var moviesessions []*model.MovieSession
	//fmt.Println(" GetMovieSessionByMovieId1.2")
	for row.Next() {
		moviesession := &model.MovieSession{}
		//fmt.Println(" GetMovieSessionByMovieId2")
		_ = row.Scan(&moviesession.ShowMovie, &moviesession.ShowCinema, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
		moviesession.State = moviesession.JudgeState()
		moviesession.Remaining = moviesession.Count()
		//fmt.Println(moviesession)
		moviesessions = append(moviesessions, moviesession)
	}
	//fmt.Println(" GetMovieSessionByMovieId3")
	page := &model.Page{
		ShowSession: moviesessions,
	}
	//fmt.Println(" GetMovieSessionByMovieId4")
	return page, nil
}

func GetMovieSessionByMovieName(moviename string) (*model.Page, error) {
	sql0 := "select movieid from movie where moviename = ?"
	row0 := utils.Db.QueryRow(sql0, moviename)

	fmt.Println("GetMovieSessionByMovieName1", row0)

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
	sql0 := "select cinemaid from cinema where cinemaname = ?"
	row0 := utils.Db.QueryRow(sql0, cinema)

	sql := "select * from moviesession where cinemaid=? and screenroom=? "
	rows, _ := utils.Db.Query(sql, row0, screenroom)
	var moviesessions []*model.MovieSession
	for rows.Next() {
		moviesession := &model.MovieSession{}
		rows.Scan(&moviesession.MovieSessionId, &moviesession.MovieId, &moviesession.CinemaId, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
		moviesessions = append(moviesessions, moviesession)
	}
	return moviesessions, nil
}
