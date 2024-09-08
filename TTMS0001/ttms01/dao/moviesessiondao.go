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

//func GetMovieSessionByMovieId(movieId string) (*model.Page, error) {
//	sql := "select movie.moviename,cinema.cinemaname,moviesession.showtime,moviesession.showinfo,moviesession.state,moviesession.screenroom,moviesession.price from moviesession join movie on movie.movieid=moviesession.movieid join cinema on cinema.cinemaid=moviesession.cinemaid where moviesession.movieid=?"
//	row, _ := utils.Db.Query(sql, movieId)
//	var moviesessions []*model.MovieSession
//	for row.Next() {
//		moviesession := &model.MovieSession{}
//		_ = row.Scan(&moviesession.ShowMovie, &moviesession.ShowCinema, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
//		moviesession.State = moviesession.JudgeState()
//		moviesession.Remaining = moviesession.Count()
//		moviesessions = append(moviesessions, moviesession)
//	}
//
//	page := &model.Page{
//		ShowSession: moviesessions,
//	}
//	return page, nil
//}

func GetMovieSessionByMovieId(movieId string) (*model.Page, error) {
	fmt.Println("GetMovieSessionByMovieId", movieId)
	sql := "SELECT movie.moviename, cinema.cinemaname, moviesession.showtime, moviesession.showinfo, moviesession.state, moviesession.screenroom, moviesession.price FROM moviesession JOIN movie ON movie.movieid = moviesession.movieid JOIN cinema ON cinema.cinemaid = moviesession.cinemaid WHERE movie.movieid = ?"
	rows, err := utils.Db.Query(sql, movieId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close() // 确保在函数退出时关闭 rows

	var moviesessions []*model.MovieSession
	for rows.Next() {
		moviesession := &model.MovieSession{}
		if err := rows.Scan(&moviesession.ShowMovie, &moviesession.ShowCinema, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		moviesession.State = moviesession.JudgeState()
		moviesession.Remaining = moviesession.Count()
		moviesessions = append(moviesessions, moviesession)
	}

	// 检查循环中的错误
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	page := &model.Page{
		ShowSession: moviesessions,
	}
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

//func SaveMovieSession(moviesession *model.MovieSession) error {
//	sql0 := "select movieid from movie where moviename=?"
//	sql1 := "select cinemaid from cinema where cinemaid=?"
//
//	sql := "insert into moviesession(movieid,cinemaid,showtime,showinfo,state,screenroom,price) values (?,?,?,?,?,?,?)"
//	fmt.Println(moviesession.MovieId, moviesession.CinemaId, moviesession.ShowTime, moviesession.ShowInfo, moviesession.State, moviesession.ShowScreen, moviesession.Price)
//	utils.Db.Exec(sql, moviesession.MovieId, moviesession.CinemaId, moviesession.ShowTime, moviesession.ShowInfo, moviesession.State, moviesession.ShowScreen, moviesession.Price)
//	return nil
//}

func SaveMovieSession(moviesession *model.MovieSession) error {
	// 检查电影是否存在
	sql0 := "SELECT movieid FROM movie WHERE moviename = ?"
	row0 := utils.Db.QueryRow(sql0, moviesession.ShowMovie) // 确保 moviesession 中有 MovieName 字段
	var movieID int
	if err := row0.Scan(&movieID); err != nil {
		return fmt.Errorf("failed to find movie: %w", err)
	}

	// 检查影院是否存在
	sql1 := "SELECT cinemaid FROM cinema WHERE cinemaname = ?"
	row1 := utils.Db.QueryRow(sql1, moviesession.ShowCinema)
	var cinemaID int
	if err := row1.Scan(&cinemaID); err != nil {
		return fmt.Errorf("failed to find cinema: %w", err)
	}

	// 插入电影场次
	fmt.Println("SaveMovieSession", moviesession.MovieId, moviesession.CinemaId, moviesession.ShowTime, moviesession.ShowInfo, moviesession.State, moviesession.ShowScreen, moviesession.Price)
	sql := "INSERT INTO moviesession (movieid, cinemaid, showtime, showinfo, state, screenroom, price) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := utils.Db.Exec(sql, movieID, cinemaID, moviesession.ShowTime, moviesession.ShowInfo, moviesession.State, moviesession.ShowScreen, moviesession.Price)
	if err != nil {
		return fmt.Errorf("failed to insert movie session: %w", err)
	}

	return nil
}

//func GetMoviesessionByCinemaAndScreen(cinema string, screenroom string) ([]*model.MovieSession, error) {
//	sql0 := "select cinemaid from cinema where cinemaname = ?"
//	row0 := utils.Db.QueryRow(sql0, cinema)
//
//	sql := "select * from moviesession where cinemaid=? and screenroom=? "
//	rows, _ := utils.Db.Query(sql, row0, screenroom)
//	var moviesessions []*model.MovieSession
//	for rows.Next() {
//		moviesession := &model.MovieSession{}
//		rows.Scan(&moviesession.MovieSessionId, &moviesession.MovieId, &moviesession.CinemaId, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
//		moviesessions = append(moviesessions, moviesession)
//	}
//	return moviesessions, nil
//}

func GetMoviesessionByCinemaAndScreen(cinema string, screenroom string) ([]*model.MovieSession, error) {
	// 查询 cinemaid
	sql0 := "SELECT cinemaid FROM cinema WHERE cinemaname = ?"
	row0 := utils.Db.QueryRow(sql0, cinema)

	var cinemaID int
	err := row0.Scan(&cinemaID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cinemaid: %w", err)
	}

	// 查询电影场次
	sql := "SELECT * FROM moviesession WHERE cinemaid = ? AND screenroom = ?"
	rows, err := utils.Db.Query(sql, cinemaID, screenroom)
	if err != nil {
		return nil, fmt.Errorf("failed to query moviesessions: %w", err)
	}
	defer rows.Close()

	var moviesessions []*model.MovieSession
	for rows.Next() {
		moviesession := &model.MovieSession{}
		err := rows.Scan(&moviesession.MovieSessionId, &moviesession.MovieId, &moviesession.CinemaId,
			&moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State,
			&moviesession.ShowScreen, &moviesession.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		moviesessions = append(moviesessions, moviesession)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return moviesessions, nil
}
