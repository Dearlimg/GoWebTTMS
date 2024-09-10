package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"ttms01/dao"
	"ttms01/model"
	"ttms01/utils"
)

func BuyTicket(w http.ResponseWriter, r *http.Request) {
	movieinfo := &model.Movie{}

	movieinfo.MovieId = r.FormValue("MovieId")
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
	movieinfo.Introduction = r.FormValue("Introduction")

	fmt.Println(movieinfo)

	page, _ := dao.GetMovieSessionByMovieId(movieinfo.MovieId)

	fmt.Println("BuyTicket1", movieinfo.MovieId)

	page.Movie = movieinfo
	//page.IsLogin = true
	//fmt.Println(page)

	flag, session := dao.IsLogin(r)
	//fmt.Println(flag)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	} else {
		page.IsLogin = false
	}
	fmt.Println("BuyTicket2")
	if dao.IsAdmin(page.Username) {
		page.IsAdmin = true
	} else {
		page.IsAdmin = false
	}

	introduction := dao.GetIntroductionByMovieName(movieinfo.MovieName)
	//fmt.Println(introduction, movieinfo.MovieName)
	//fmt.Println(introduction)
	page.Introductions = introduction
	fmt.Println("BuyTicket3")
	comments := dao.GetCommentsByMovieName(movieinfo.MovieId)
	//fmt.Println(comments)
	page.Comments = comments
	//fmt.Println(page.ShowSession)
	fmt.Println("BuyTicket4")
	t := template.Must(template.ParseFiles("views/pages/trade/showinfo.html"))
	t.Execute(w, page)
}

func Buy(w http.ResponseWriter, r *http.Request) {
	movie_session := &model.MovieSession{}
	movie_session.ShowCinema = r.FormValue("ShowCinema")
	movie_session.ShowScreen = r.FormValue("ShowScreen")
	movie_session.ShowTime = r.FormValue("ShowTime")
	movie_session.ShowMovie = r.FormValue("ShowMovie")
	movie_session.ShowInfo = r.FormValue("ShowInfo")
	movie_session.Price, _ = strconv.ParseFloat(r.FormValue("Price"), 64)

	nums, _ := dao.ParseInfo(movie_session.ShowInfo)
	//fmt.Println(nums)
	//page, _ := dao.GetPageMovieByCinemaName(movie_session.ShowMovie)
	movie, _ := dao.GetMovieInfoByMovieName(movie_session.ShowMovie)
	page := &model.Page{}
	page.ShowSession1 = movie_session
	page.Showinfo = nums
	page.Movie = movie
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

	t := template.Must(template.ParseFiles("views/pages/trade/buy.html"))
	t.Execute(w, page)
}

func Bill(w http.ResponseWriter, r *http.Request) {
	finstr := ""
	for i := 0; i < 20; i++ {
		istr := strconv.Itoa(i)
		flag := r.FormValue("seat" + istr)
		fmt.Println("flag", flag)
		if flag == "0" {
			flag = "1"
		} else if flag == "1" {
			flag = "1"
		} else {
			flag = "0"
		}
		finstr = finstr + flag
	}

	fmt.Println(finstr)

	Price1, _ := strconv.ParseFloat(r.FormValue("Price"), 64)

	moviesession := &model.MovieSession{
		ShowCinema: r.FormValue("ShowCinema"),
		ShowScreen: r.FormValue("ShowScreen"),
		ShowTime:   r.FormValue("ShowTime"),
		ShowMovie:  r.FormValue("ShowMovie"),
		ShowInfo:   r.FormValue("ShowInfo"),
		Price:      Price1,
	}

	fmt.Println("Bill1", moviesession.ShowCinema)

	_ = dao.SaveData(finstr, moviesession)

	nums := dao.Compare(finstr, moviesession.ShowInfo)

	var count int = len(nums)
	floatcount := float64(count)
	sum := floatcount * Price1
	_, session := dao.IsLogin(r)

	page := &model.Page{
		Seatinfo:     nums,
		Session:      session,
		ShowSession1: moviesession,
		SumPrice:     sum,
	}

	_, session = dao.IsLogin(r)

	packed := utils.PackTicketData(nums)
	realseat := utils.ParseTicketData(packed)
	fmt.Println("Bill", realseat)
	buyer := r.FormValue("owner")

	ticket := &model.Ticket{
		Owner:  buyer,
		Movie:  moviesession.ShowMovie,
		Screen: moviesession.ShowScreen,
		Cinema: moviesession.ShowCinema,
		Time:   moviesession.ShowTime,
		Seat:   realseat,
		Price:  Price1,
	}

	dao.AddTicket(ticket)

	t := template.Must(template.ParseFiles("views/pages/trade/bill.html"))
	t.Execute(w, page)
}
func ShowTickets(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	fmt.Println("showtickets1")
	page := &model.Page{
		IsLogin:  true,
		Username: session.UserName,
	}
	if dao.IsAdmin(page.Username) {
		fmt.Println("showtickets2")
		page.IsAdmin = true
		tickets := dao.GetAllTickets()

		page.Tickets = tickets
		t := template.Must(template.ParseFiles("views/pages/user/account.html"))
		t.Execute(w, page)
	} else {
		page.IsAdmin = false
		fmt.Println("ShowTickets", session.UserName)
		tickets, _ := dao.GetTicketsByName(session.UserName)
		page.Tickets = tickets

		t := template.Must(template.ParseFiles("views/pages/user/account.html"))
		t.Execute(w, page)
	}
}

func ReturnTicket(w http.ResponseWriter, r *http.Request) {
	ticket := &model.Ticket{
		Owner:  r.FormValue("Owner"),
		Movie:  r.FormValue("Movie"),
		Cinema: r.FormValue("Cinema"),
		Screen: r.FormValue("Screen"),
		Seat:   r.FormValue("Seat"),
		Time:   r.FormValue("Time"),
	}
	price, _ := strconv.ParseFloat(r.FormValue("Price"), 64)
	ticket.Price = price

	dao.DeleteTicketByAllInfo(ticket)

	ShowTickets(w, r)
}

func ModifyTicket(w http.ResponseWriter, r *http.Request) {
	ticket := &model.Ticket{
		Owner:  r.FormValue("Owner"),
		Movie:  r.FormValue("Movie"),
		Cinema: r.FormValue("Cinema"),
		Screen: r.FormValue("Screen"),
		Seat:   r.FormValue("Seat"),
		Time:   r.FormValue("Time"),
	}
	price, _ := strconv.ParseFloat(r.FormValue("Price"), 64)
	ticket.Price = price
	//fmt.Println(price)
	t := template.Must(template.ParseFiles("views/pages/admin/modifyticket.html"))
	t.Execute(w, ticket)
}

func ModifyTicket1(w http.ResponseWriter, r *http.Request) {
	ticket := &model.Ticket{
		Owner:  r.FormValue("owner"),
		Movie:  r.FormValue("movie"),
		Cinema: r.FormValue("cinema"),
		Screen: r.FormValue("hall"),
		Seat:   r.FormValue("seat"),
		Time:   r.FormValue("time"),
	}
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	ticket.Price = price

	//fmt.Println("1", ticket)

	dao.DeleteTicketWithoutSeat(ticket)
	dao.AddTicket(ticket)

	newseat := strings.Replace(ticket.Seat, " ", "_", -1)
	//fmt.Println("2 new", newseat)
	moviesession, _ := dao.GetMovieSessionByTicket(ticket)
	//fmt.Println("3", moviesession.ShowInfo)

	oldseats := utils.SeatsToNumbers(moviesession.ShowInfo)
	//fmt.Println("6 old", oldseats)

	if len(oldseats) > len(newseat) {
		//res := strings.Replace(oldseats, newseat, "", 1)
		//modifedinfo := dao.ModifySessionInfo(moviesession.ShowInfo, res, "sell")
		//fmt.Println(modifedinfo)
		//dao.ModifyShowSessionSeat(moviesession, modifedinfo)
	} else {
		//res := strings.Replace(newseat, oldseats, "", 1)
		//modifedinfo := dao.ModifySessionInfo(moviesession.ShowInfo, res, "buy")
		//fmt.Println(modifedinfo)
		//dao.ModifyShowSessionSeat(moviesession, modifedinfo)
	}
	GetPageMovie(w, r)
}
