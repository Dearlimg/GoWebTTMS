package main

import (
	"html/template"
	"net/http"
	"ttms01/controller"
	"ttms01/dao"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, session)
}

func main() {
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))

	http.HandleFunc("/main", controller.GetPageMovie)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/movieSearch", controller.MovieSearch)
	http.HandleFunc("/cinemaSearch", controller.CinemaSearch)
	http.HandleFunc("/exit", controller.Exit)
	http.HandleFunc("/buyticket", controller.BuyTicket)
	http.HandleFunc("/choicescreenroom", controller.ChoiceScreenRoom)
	http.HandleFunc("/buy", controller.Buy)
	http.HandleFunc("/bill", controller.Bill)
	http.HandleFunc("/showtickets", controller.ShowTickets)
	http.HandleFunc("/jumptoregister", controller.JumpToRegister)
	http.HandleFunc("/jumptologin", controller.JumpToLogin)
	http.HandleFunc("/returnticket", controller.ReturnTicket)
	http.HandleFunc("/backofficemanagement", controller.BackOfficeManagement)
	http.HandleFunc("/mainsearch", controller.MainSearch)

	http.HandleFunc("/adminmodifymovie", controller.AdminModifyMovie)
	http.HandleFunc("/deletemoviesession", controller.DeleteMovieSession)
	http.HandleFunc("/addmoviesession", controller.AddMovieSession)
	http.HandleFunc("/addmoviesession1", controller.AddMovieSession1)
	http.HandleFunc("/modifyticket", controller.ModifyTicket)
	http.HandleFunc("/modifyticket1", controller.ModifyTicket1)
	http.HandleFunc("/addmovie", controller.AddMovie)
	http.HandleFunc("/addmovie1", controller.AddMovie1)
	http.HandleFunc("/deletemovie", controller.DeleteMovie)
	http.HandleFunc("/modifymovie", controller.ModifyMovie)
	http.HandleFunc("/modifymovie1", controller.ModifyMovie1)
	http.HandleFunc("/addcinema", controller.AddCinema)
	http.HandleFunc("/addcinema1", controller.AddCinema1)
	http.HandleFunc("/deletecinema", controller.DeleteCinema)
	http.HandleFunc("/modifycinema", controller.ModifyCinema)
	http.HandleFunc("/modifycinema1", controller.ModifyCinema1)

	http.HandleFunc("/submitcomment", controller.SubmitComment)
	http.HandleFunc("/deletecomment", controller.DeleteComment)

	http.ListenAndServe(":8001", nil)
}
