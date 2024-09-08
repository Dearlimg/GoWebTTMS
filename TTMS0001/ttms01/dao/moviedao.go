package dao

import (
	"fmt"
	"strconv"
	"ttms01/model"
	"ttms01/utils"
)

func GetPageMovie(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sql := "select count(*) from movie"
	var totalRecord int64
	row := utils.Db.QueryRow(sql)
	row.Scan(&totalRecord)

	var pageSize int64 = 6
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	sql1 := "select * from movie where state>0 order by score desc limit ?,?"
	rows, err := utils.Db.Query(sql1, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		//fmt.Println(movie)
		movies = append(movies, movie)
	}
	page := &model.Page{
		Movies:      movies,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetAllMovies() (*model.Page, error) {
	sql1 := "select count(*) from movie"
	var totalRecord int64
	row := utils.Db.QueryRow(sql1)
	row.Scan(&totalRecord)

	sql := `select * from movie where state>0`
	var movies []*model.Movie
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		fmt.Println("GetAllMovies", movie)
		movies = append(movies, movie)
	}
	page := &model.Page{
		Movies:      movies,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetPageMovieSessionByCinemaName(cinema string) (*model.Page, error) {

	sql1 := "select cinemaid from cinema where cinemaname=?"
	row := utils.Db.QueryRow(sql1, cinema)

	sql := "select * from moviesession where cinemaid=?"
	var moviesessions []*model.MovieSession
	rows, err := utils.Db.Query(sql, row)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		moviesession := &model.MovieSession{}
		rows.Scan(&moviesession.MovieSessionId, &moviesession.MovieId, &moviesession.CinemaId, &moviesession.ShowTime, &moviesession.ShowInfo, &moviesession.State, &moviesession.ShowScreen, &moviesession.Price)
		moviesession.State = moviesession.JudgeState()
		moviesession.Remaining = moviesession.Count()
		moviesession.MovieImgPath, _ = GetMovieImgByMovieName(moviesession.ShowMovie)
		moviesessions = append(moviesessions, moviesession)
	}
	page := &model.Page{
		ShowSession: moviesessions,
	}
	return page, nil
}

func GetMovieInfoByMovieName(movieName string) (*model.Movie, error) {
	sql := "select * from movie where moviename=?"
	row := utils.Db.QueryRow(sql, movieName)
	movie := &model.Movie{}
	row.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
	return movie, nil
}

func GetMovieImgByMovieName(movieName string) (string, error) {
	sql := "select imgpath from movie where moviename=?"
	row := utils.Db.QueryRow(sql, movieName)
	var imgPath string
	row.Scan(&imgPath)
	return imgPath, nil
}

func GetComingMovies(start string, end string, now string) []*model.Movie {
	sql1 := "select count(*) from movie"
	var totalRecord int
	row := utils.Db.QueryRow(sql1)
	row.Scan(&totalRecord)
	if start == "" {
		start = "0"
	}
	if end == "" {
		end = strconv.Itoa(totalRecord)
	}

	//fmt.Println(start, end, totalRecord, now)

	sql := "SELECT * FROM movie WHERE showtime > ? and state>0 ORDER BY showtime ASC LIMIT ?, ?"
	rows, err := utils.Db.Query(sql, now, start, end)
	if err != nil {
		return nil
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		fmt.Println(movie)
		movies = append(movies, movie)
	}
	return movies
}

func GetHotMovies(start string, end string, now string) []*model.Movie {
	sql1 := "select count(*) from movie"
	var totalRecord int
	row := utils.Db.QueryRow(sql1)
	row.Scan(&totalRecord)
	if start == "" {
		start = "0"
	}
	if end == "" {
		end = strconv.Itoa(totalRecord)
	}

	fmt.Println(start, end, totalRecord, now)

	sql := "SELECT * FROM movie WHERE showtime < ? and state>0 ORDER BY boxoffice DESC LIMIT ?, ?"
	rows, err := utils.Db.Query(sql, now, start, end)
	if err != nil {
		return nil
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		fmt.Println(movie)
		movies = append(movies, movie)
	}
	return movies
}

func GetClassicMovies(start string, end string, now string) []*model.Movie {
	sql1 := "select count(*) from movie"
	var totalRecord int
	row := utils.Db.QueryRow(sql1)
	row.Scan(&totalRecord)
	if start == "" {
		start = "0"
	}
	if end == "" {
		end = strconv.Itoa(totalRecord)
	}

	//fmt.Println(start, end, totalRecord, now)

	sql := "SELECT * FROM movie WHERE showtime < ? and state>0 ORDER BY score DESC LIMIT ?, ?"
	rows, err := utils.Db.Query(sql, now, start, end)
	if err != nil {
		return nil
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		fmt.Println(movie)
		movies = append(movies, movie)
	}
	return movies
}

func GetMoviesByCondition(genre, region, decade, sort string) []*model.Movie {
	sql := "SELECT * FROM movie WHERE 1=1"
	var args []interface{}

	if genre != "" {
		sql += " AND genre=?"
		args = append(args, genre)
	}
	if region != "" {
		sql += " AND area=?"
		args = append(args, region)
	}

	// 确定排序列
	if sort != "" {
		sql += " ORDER BY score " + sort
	}
	//fmt.Println(sql)

	// 执行查询
	rows, _ := utils.Db.Query(sql, args...)

	defer rows.Close()

	movies := []*model.Movie{}
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		movies = append(movies, movie)
	}
	return movies
}

func GetMovieByKeyWord(keyword string) []*model.Movie {
	sql := "SELECT * FROM movie WHERE moviename like " + "'%" + keyword + "%'" + "and state>0"
	//fmt.Println(sql)
	rows, _ := utils.Db.Query(sql)
	movies := []*model.Movie{}
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieId, &movie.MovieName, &movie.ActorName, &movie.Showtime, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.ImgPath, &movie.Duration, &movie.Introduction, &movie.State)
		movies = append(movies, movie)
	}
	return movies
}

func SaveMovie(movie *model.Movie) {
	sql := "insert into movie(movieid,moviename,actorname,showtime,score,boxoffice,genre,area,imgpath,Duration,introduction,state) values (?,?,?,?,?,?,?,?,?,?,?,?)"
	utils.Db.Exec(sql, movie.MovieId, movie.MovieName, movie.ActorName, movie.Showtime, movie.Score, movie.BoxOffice, movie.Genre, movie.Area, movie.ImgPath, movie.Duration, movie.Introduction, movie.State)
}

func DeleteMovieByMovieName(moviename string) {
	sql := "update movie set state=0 where moviename=?"
	utils.Db.Exec(sql, moviename)
}

//func UpdateMovieByMovieName(movie *model.Movie, moviename string) {
//	sql := "update movie set movieid=? moviename=? , actorname=? , showtime=? , score=? , boxoffice=? , genre=? , area=? ,imgpath=? , duration=? , introduction where moviename = ?"
//	utils.Db.Exec(sql, movie.MovieId, movie.MovieName, movie.ActorName, movie.Showtime, movie.Score, movie.BoxOffice, movie.Genre, movie.Area, movie.ImgPath, movie.Duration, movie.Introduction, movie.State)
//}

func GetIntroductionByMovieName(moviename string) *model.Introduction {
	sql := "select introduction from movie where moviename = ?"
	row := utils.Db.QueryRow(sql, moviename)
	res := &model.Introduction{}
	row.Scan(&res.Intro)
	return res
}
