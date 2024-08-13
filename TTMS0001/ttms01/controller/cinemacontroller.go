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
	//fmt.Println(district, rank)

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
		page, _ := dao.GetAllCinema()
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

	//

	//

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

	Price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	//fmt.Println("测试1", moviesession.ShowCinema, moviesession.ShowScreen)
	allmoviesession, _ := dao.GetMoviesessionByCinemaAndScreen(moviesession.ShowCinema, moviesession.ShowScreen)
	flag := -1
	for _, v := range allmoviesession {
		//res, _ := utils.IsWithinRange(moviesession.ShowTime, v.ShowTime, v.GetMovieDuration(moviesession.ShowMovie))
		res, _ := utils.IsWithinRange(moviesession.ShowTime, v.ShowTime, "180")
		fmt.Println("AddMovieSession1:", moviesession.ShowTime, v.ShowTime, v.GetMovieDuration(moviesession.ShowMovie), res)
		if res == true {
			flag = 1
		}
	}
	if flag == -1 {
		moviesession.Price = Price
		moviesession.Remaining = moviesession.Count()
		//fmt.Println(moviesession)
		dao.SaveMovieSession(moviesession)
		GetPageMovie(w, r)
	} else {
		moviename := moviesession.ShowMovie

		//

		//

		moviesession := &model.MovieSession{
			ShowMovie:  moviename,
			ShowCinema: moviesession.ShowCinema,
			ShowScreen: moviesession.ShowScreen,
			Price:      Price,
		}
		page := &model.Page{
			ShowSession1: moviesession,
		}
		page.Message = "时间冲突"
		t := template.Must(template.ParseFiles("views/pages/admin/addmoviesession.html"))
		t.Execute(w, page)
	}
}

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
