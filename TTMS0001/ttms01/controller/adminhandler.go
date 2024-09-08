package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"ttms01/dao"
	"ttms01/model"
)

func BackOfficeManagement(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/admin/management.html"))
	t.Execute(w, nil)
}

func AdminModifyMovie(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/admin/modifymovie.html"))
	t.Execute(w, nil)
}

func DeleteMovieSession(w http.ResponseWriter, r *http.Request) {
	movie_session := &model.MovieSession{}
	movie_session.ShowCinema = r.FormValue("ShowCinema")
	movie_session.ShowScreen = r.FormValue("ShowScreen")
	movie_session.ShowTime = r.FormValue("ShowTime")
	movie_session.ShowMovie = r.FormValue("ShowMovie")
	movie_session.ShowInfo = r.FormValue("ShowInfo")
	movie_session.Price, _ = strconv.ParseFloat(r.FormValue("Price"), 64)

	dao.DeleteMovieSession(movie_session)

	//测试一下
	GetPageMovie(w, r)
}
