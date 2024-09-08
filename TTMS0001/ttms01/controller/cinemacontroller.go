package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"ttms01/dao"
	"ttms01/model"
	"ttms01/utils"
)

func CinemaSearch(w http.ResponseWriter, r *http.Request) {
	district := r.FormValue("district")
	rank := r.FormValue("rank")
	if district != "" || rank != "" {
		page := &model.Page{
			Cinema: dao.GetCinemaByCondition(district, rank),
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
		//fmt.Println(page.Cinema[0].CinemaName, "dwawda")
		t := template.Must(template.ParseFiles("views/pages/Home/cinema.html"))
		t.Execute(w, page)
	} else {
		//fmt.Println("cinemasearch2")
		page, _ := dao.GetAllCinema()
		flag, session := dao.IsLogin(r)
		//fmt.Println("cinemasearch2.1")
		if flag {
			page.IsLogin = true
			page.Username = session.UserName
		}
		//fmt.Println("cinemasearch2.2", page)
		//fmt.Println("cinemasearch3", page.Username)
		if page.Username != "" {
			//fmt.Println("cinemasearch3.1")
			if dao.IsAdmin(page.Username) {
				page.IsAdmin = true
			} else {
				page.IsAdmin = false
			}
		}

		//fmt.Println("cinemasearch4")
		//fmt.Println(page.Cinema[0].CinemaName, "dwawda")
		t := template.Must(template.ParseFiles("views/pages/Home/cinema.html"))
		t.Execute(w, page)
	}
}

func ChoiceScreenRoom(w http.ResponseWriter, r *http.Request) {
	cinema := &model.Cinema{}
	cinema.CinemaName = r.FormValue("CinemaName")
	cinema.Place = r.FormValue("Place")
	cinema.Rank = r.FormValue("Rank")

	//fmt.Println(cinema.CinemaName)

	page, _ := dao.GetPageMovieSessionByCinemaName(cinema.CinemaName)
	//得到目标影院所安拍播放的场次
	//fmt.Println(page)
	showsession := &model.MovieSession{
		ShowCinema: cinema.CinemaName,
	}
	page.ShowSession1 = showsession
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

	t := template.Must(template.ParseFiles("views/pages/cinema/screenroom.html"))
	t.Execute(w, page)
}

func AddMovieSession(w http.ResponseWriter, r *http.Request) {
	moviename := r.FormValue("Moviename")
	moviesession := &model.MovieSession{
		ShowMovie:  moviename,
		ShowCinema: r.FormValue("cinemaname"),
	}
	page := &model.Page{
		ShowSession1: moviesession,
	}

	t := template.Must(template.ParseFiles("views/pages/admin/addmoviesession.html"))
	t.Execute(w, page)
}

func AddMovieSession1(w http.ResponseWriter, r *http.Request) {
	moviesession := &model.MovieSession{
		ShowMovie:  r.FormValue("movie"),
		ShowCinema: r.FormValue("cinema"),
		ShowScreen: r.FormValue("hall"),
		ShowTime:   r.FormValue("time"),
		ShowInfo:   "00000000000000000000",
	}

	Price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}
	moviesession.Price = Price

	// Query to get cinemaid
	var cinemaId string
	sql01 := "SELECT cinemaid FROM cinema WHERE cinemaname = ?"
	err = utils.Db.QueryRow(sql01, moviesession.ShowCinema).Scan(&cinemaId)
	if err != nil {
		http.Error(w, "Cinema not found", http.StatusNotFound)
		return
	}
	moviesession.CinemaId = cinemaId

	// Check for conflicting movie sessions
	allmoviesession, err := dao.GetMoviesessionByCinemaAndScreen(moviesession.ShowCinema, moviesession.ShowScreen)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	flag := -1
	for _, v := range allmoviesession {
		res, _ := utils.IsWithinRange(moviesession.ShowTime, v.ShowTime, "180")
		if res {
			flag = 1
			break
		}
	}

	// Get movieId
	var movieId string
	err = utils.Db.QueryRow("SELECT movieid FROM movie WHERE moviename = ?", moviesession.ShowMovie).Scan(&movieId)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	moviesession.MovieId = movieId

	// Save movie session or return conflict message
	if flag == -1 {
		moviesession.Remaining = moviesession.Count()
		if err := dao.SaveMovieSession(moviesession); err != nil {
			http.Error(w, "Failed to save session", http.StatusInternalServerError)
			return
		}
		GetPageMovie(w, r)
	} else {
		page := &model.Page{
			ShowSession1: &model.MovieSession{
				ShowMovie:  moviesession.ShowMovie,
				ShowCinema: moviesession.ShowCinema,
				ShowScreen: moviesession.ShowScreen,
				Price:      Price,
			},
			Message: "时间冲突",
		}
		t := template.Must(template.ParseFiles("views/pages/admin/addmoviesession.html"))
		if err := t.Execute(w, page); err != nil {
			http.Error(w, "Template execution failed", http.StatusInternalServerError)
		}
	}
}

//func AddMovieSession1(w http.ResponseWriter, r *http.Request) {
//	moviesession := &model.MovieSession{
//		ShowMovie:  r.FormValue("movie"),
//		ShowCinema: r.FormValue("cinema"),
//		ShowScreen: r.FormValue("hall"),
//		ShowTime:   r.FormValue("time"),
//		ShowInfo:   "00000000000000000000",
//	}
//	fmt.Println("AddMovieSession0")
//	Price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
//	//fmt.Println("测试1", moviesession.ShowCinema, moviesession.ShowScreen)
//	fmt.Println("AddMovieSession0.1")
//	allmoviesession, _ := dao.GetMoviesessionByCinemaAndScreen(moviesession.ShowCinema, moviesession.ShowScreen)
//
//	fmt.Println("AddMovieSession1")
//	flag := -1
//	for _, v := range allmoviesession {
//		//res, _ := utils.IsWithinRange(moviesession.ShowTime, v.ShowTime, v.GetMovieDuration(moviesession.ShowMovie))
//		res, _ := utils.IsWithinRange(moviesession.ShowTime, v.ShowTime, "180")
//		fmt.Println("AddMovieSession1:", moviesession.ShowTime, v.ShowTime, v.GetMovieDuration(moviesession.ShowMovie), res)
//		if res == true {
//			flag = 1
//		}
//	}
//
//	sql0 := "select movieid from movie where moviename=?"
//	utils.Db.QueryRow(sql0, moviesession.ShowMovie)
//
//	fmt.Println("AddMovieSession2")
//	if flag == -1 {
//		moviesession.Price = Price
//		moviesession.Remaining = moviesession.Count()
//		fmt.Println("adwadawdawdawwadaw", moviesession.MovieId)
//		dao.SaveMovieSession(moviesession)
//		fmt.Println("AddMovieSession3")
//		GetPageMovie(w, r)
//	} else {
//		moviename := moviesession.ShowMovie
//
//		//
//
//		//
//
//		moviesession := &model.MovieSession{
//			ShowMovie:  moviename,
//			ShowCinema: moviesession.ShowCinema,
//			ShowScreen: moviesession.ShowScreen,
//			Price:      Price,
//		}
//		fmt.Println("AddMovieSession4")
//		page := &model.Page{
//			ShowSession1: moviesession,
//		}
//		page.Message = "时间冲突"
//		t := template.Must(template.ParseFiles("views/pages/admin/addmoviesession.html"))
//		t.Execute(w, page)
//	}
//}

func AddCinema(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/admin/addcinema.html"))
	t.Execute(w, nil)
}

func AddCinema1(w http.ResponseWriter, r *http.Request) {
	cinema := &model.Cinema{
		CinemaName: r.FormValue("cinema_name"),
		Place:      r.FormValue("place"),
		Rank:       r.FormValue("cinema_rank"),
	}
	fmt.Println(cinema)
	dao.AddCinema(cinema)
	CinemaSearch(w, r)
}

func DeleteCinema(w http.ResponseWriter, r *http.Request) {
	//cinema:=&model.Cinema{
	CinemaName := r.FormValue("cinemaname")
	//Place:      r.FormValue("place"),
	//Rank:       r.FormValue("cinema_rank"),
	//}
	dao.DeleteCinemaByCinemaName(CinemaName)
	CinemaSearch(w, r)
}

func ModifyCinema(w http.ResponseWriter, r *http.Request) {
	cinema := &model.Cinema{
		CinemaName: r.FormValue("cinemaname"),
		Place:      r.FormValue("place"),
		Rank:       r.FormValue("cinemarank"),
	}
	t := template.Must(template.ParseFiles("views/pages/admin/modifycinema.html"))
	t.Execute(w, cinema)
}

func ModifyCinema1(w http.ResponseWriter, r *http.Request) {
	cinema := &model.Cinema{
		CinemaName: r.FormValue("cinemaname"),
		Place:      r.FormValue("place"),
		Rank:       r.FormValue("cinemarank"),
	}
	//fmt.Println(cinema)
	dao.SaveCinema(cinema, cinema.CinemaName)
	GetPageMovie(w, r)
}
