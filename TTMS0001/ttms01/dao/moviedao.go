package dao

import (
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
	sql1 := "select * from movie order by score desc limit ?,?"
	rows, err := utils.Db.Query(sql1, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
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

	sql := `select * from movie`
	var movies []*model.Movie
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
		movies = append(movies, movie)
	}
	page := &model.Page{
		Movies:      movies,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetPageMovieSessionByCinemaName(cinema string) (*model.Page, error) {
	sql := "select * from movie_session where cinema=?"
	var moviesessions []*model.MovieSession
	rows, err := utils.Db.Query(sql, cinema)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		moviesession := &model.MovieSession{}
		rows.Scan(&moviesession.ShowCinema, &moviesession.ShowScreen, &moviesession.ShowTime, &moviesession.ShowMovie, &moviesession.ShowInfo, &moviesession.Price)
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
	row.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
	return movie, nil
}

func GetMovieImgByMovieName(movieName string) (string, error) {
	sql := "select img_path from movie where moviename=?"
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

	sql := "SELECT * FROM movie WHERE showtime > ? ORDER BY showtime ASC LIMIT ?, ?"
	rows, err := utils.Db.Query(sql, now, start, end)
	if err != nil {
		return nil
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
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

	//fmt.Println(start, end, totalRecord, now)

	sql := "SELECT * FROM movie WHERE showtime < ? ORDER BY boxoffice DESC LIMIT ?, ?"
	rows, err := utils.Db.Query(sql, now, start, end)
	if err != nil {
		return nil
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
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

	sql := "SELECT * FROM movie WHERE showtime < ? ORDER BY score DESC LIMIT ?, ?"
	rows, err := utils.Db.Query(sql, now, start, end)
	if err != nil {
		return nil
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
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
	if decade != "" {
		sql += " AND time=?"
		args = append(args, decade)
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
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
		movies = append(movies, movie)
	}
	return movies
}

func GetMovieByKeyWord(keyword string) []*model.Movie {
	sql := "SELECT * FROM movie WHERE moviename like " + "'%" + keyword + "%'"
	//fmt.Println(sql)
	rows, _ := utils.Db.Query(sql)
	movies := []*model.Movie{}
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieName, &movie.ActorName, &movie.Time, &movie.Score, &movie.BoxOffice, &movie.Genre, &movie.Area, &movie.Age, &movie.ImgPath, &movie.Duration, &movie.Cinema, &movie.Showtime)
		movies = append(movies, movie)
	}
	return movies
}

func SaveMovie(movie *model.Movie) {
	sql := "insert into movie(moviename,actorname,time,score,boxoffice,genre,area,age,img_path,Duration,cinema,showtime) values (?,?,?,?,?,?,?,?,?,?,?,?)"
	utils.Db.Exec(sql, movie.MovieName, movie.ActorName, movie.Time, movie.Score, movie.BoxOffice, movie.Genre, movie.Area, movie.Age, movie.ImgPath, movie.Duration, movie.Cinema, movie.Showtime)
}

func DeleteMovieByMovieName(moviename string) {
	sql := "delete from movie where moviename = ?"
	utils.Db.Exec(sql, moviename)
}

func UpdateMovieByMovieName(movie *model.Movie, moviename string) {
	sql := "update movie set moviename=? , actorname=? , time=? , score=? , boxoffice=? , genre=? , area=? , age=? ,img_path=? , Duration=? , showtime=? where moviename = ?"
	utils.Db.Exec(sql, movie.MovieName, movie.ActorName, movie.Time, movie.Score, movie.BoxOffice, movie.Genre, movie.Area, movie.Age, movie.ImgPath, movie.Duration, movie.Showtime, moviename)
}

func GetIntroductionByMovieName(moviename string) *model.Introduction {
	sql := "select * from introduce where movie = ?"
	row := utils.Db.QueryRow(sql, moviename)
	res := &model.Introduction{}
	row.Scan(&res.MovieName, &res.Intro)
	return res
}
