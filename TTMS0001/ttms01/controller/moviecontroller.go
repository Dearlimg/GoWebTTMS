package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"ttms01/dao"
	"ttms01/model"
)

func GetPageMovie(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("start1")
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageMovie(pageNo)

	now := time.Now()

	format := "2006-01-02 15:04" // 24小时制，小时范围是0-23

	formattedTime := now.Format(format)

	page.ComingMovies = dao.GetComingMovies("0", "6", formattedTime)
	page.Movies = dao.GetHotMovies("", "6", formattedTime)
	page.ClassicMovies = dao.GetClassicMovies("", "6", formattedTime)
	//fmt.Println("start2")
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}

	if dao.IsAdmin(page.Username) {
		page.IsAdmin = true
	} else {
		page.IsAdmin = false
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func MovieSearch(w http.ResponseWriter, r *http.Request) {

	genre := r.FormValue("genre")
	region := r.FormValue("region")
	decade := r.FormValue("decade")
	sort := r.FormValue("sort")

	if genre != "" || region != "" || decade != "" || sort != "" {
		page := &model.Page{
			Movies: dao.GetMoviesByCondition(genre, region, decade, sort),
		}
		flag, session := dao.IsLogin(r)
		if flag {
			page.IsLogin = true
			page.Username = session.UserName
		}
		if dao.IsAdmin(page.Username) {
			page.IsAdmin = true
		} else {
			page.IsAdmin = false
		}
		t := template.Must(template.ParseFiles("views/pages/Home/movie.html"))
		t.Execute(w, page)
	} else {
		page, _ := dao.GetAllMovies()
		fmt.Println("MovieSearch1")
		flag, session := dao.IsLogin(r)
		if flag {
			page.IsLogin = true
			page.Username = session.UserName
		}
		if dao.IsAdmin(page.Username) {
			page.IsAdmin = true
		} else {
			page.IsAdmin = false
		}
		t := template.Must(template.ParseFiles("views/pages/Home/movie.html"))
		t.Execute(w, page)
	}

	//flag, session := dao.IsLogin(r)
	//if flag {
	//	page.IsLogin = true
	//	page.Username = session.UserName
	//}
	//t := template.Must(template.ParseFiles("views/pages/Home/movie.html"))
	//t.Execute(w, page)
}

func MainSearch(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("search")
	//fmt.Println(key, "dwadw")
	movies := dao.GetMovieByKeyWord(key)
	//fmt.Println(movies)
	page := &model.Page{
		Movies: movies,
	}
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}
	t := template.Must(template.ParseFiles("views/pages/trade/mainsearch.html"))
	t.Execute(w, page)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/admin/addmovie.html"))
	t.Execute(w, nil)
}

func AddMovie1(w http.ResponseWriter, r *http.Request) {
	movie := &model.Movie{
		MovieName: r.FormValue("moviename"),
		ActorName: r.FormValue("actorname"),
		Showtime:  r.FormValue("Showtime"),
		Genre:     r.FormValue("genre"),
		Area:      r.FormValue("area"),
		Age:       r.FormValue("age"),
		ImgPath:   r.FormValue("img_path"),
		Duration:  r.FormValue("Duration"),
		//Cinema:    r.FormValue("cinema"),
	}
	movie.Score, _ = strconv.ParseFloat(r.FormValue("score"), 64)
	movie.BoxOffice, _ = strconv.ParseInt(r.FormValue("boxoffice"), 10, 64)
	dao.SaveMovie(movie)
	GetPageMovie(w, r)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	moviename := r.FormValue("moviename")
	dao.DeleteMovieByMovieName(moviename)
	MovieSearch(w, r)
}

func ModifyMovie(w http.ResponseWriter, r *http.Request) {
	movieinfo := &model.Movie{}

	movieinfo.MovieName = r.FormValue("MovieName")
	movieinfo.ActorName = r.FormValue("ActorName")
	movieinfo.Showtime = r.FormValue("Showtime")
	corestr := r.FormValue("Score")
	movieinfo.Score, _ = strconv.ParseFloat(corestr, 64)
	BoxOfficeStr := r.FormValue("BoxOffice")
	movieinfo.BoxOffice, _ = strconv.ParseInt(BoxOfficeStr, 10, 64)
	movieinfo.Genre = r.FormValue("Genre")
	movieinfo.Area = r.FormValue("Area")
	movieinfo.ImgPath = r.FormValue("ImgPath")
	movieinfo.Duration = r.FormValue("Duration")
	fmt.Println(movieinfo.Showtime)
	page := &model.Page{
		Movie: movieinfo,
	}

	t := template.Must(template.ParseFiles("views/pages/admin/modifymovie.html"))
	t.Execute(w, page)
}

func ModifyMovie1(w http.ResponseWriter, r *http.Request) {
	movieinfo := &model.Movie{}
	movieinfo.MovieName = r.FormValue("moviename")
	movieinfo.ActorName = r.FormValue("actorname")
	movieinfo.Showtime = r.FormValue("Showtime")
	corestr := r.FormValue("score")
	movieinfo.Score, _ = strconv.ParseFloat(corestr, 64)
	BoxOfficeStr := r.FormValue("boxoffice")
	movieinfo.BoxOffice, _ = strconv.ParseInt(BoxOfficeStr, 10, 64)
	movieinfo.Genre = r.FormValue("genre")
	movieinfo.Area = r.FormValue("area")
	movieinfo.ImgPath = r.FormValue("img_path")
	movieinfo.Duration = r.FormValue("duration")
	movieinfo.Showtime = r.FormValue("showtime")
	fmt.Println(movieinfo)
	//dao.UpdateMovieByMovieName(movieinfo, movieinfo.MovieName)
	GetPageMovie(w, r)
}
