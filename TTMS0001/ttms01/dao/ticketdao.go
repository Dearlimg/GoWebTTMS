package dao

import (
	"strconv"
	"strings"
	"ttms01/model"
	"ttms01/utils"
)

func SaveData(data string, session *model.MovieSession) error {
	sql := "update moviesession set showinfo=? where cinemaid=? and screenroom=? and showtime=? and movieid=?"
	_, err := utils.Db.Exec(sql, data, session.CinemaId, session.ShowScreen, session.ShowTime, session.MovieId)
	if err != nil {
		return err
	}
	return nil
}

func CheckData(session *model.MovieSession) *model.MovieSession {
	sql := "select * from movie_session where cinema=? and screenroom=? and movietime=? and movie=? and info=?"
	row := utils.Db.QueryRow(sql, session.ShowCinema, session.ShowScreen, session.ShowTime, session.ShowMovie, session.ShowInfo)
	res := &model.MovieSession{}
	row.Scan(res.ShowCinema, res.ShowScreen, res.ShowTime, res.ShowMovie, res.ShowInfo)
	return res
}

func Compare(new string, old string) []int {
	res := make([]int, 0, 20)
	for i := 0; i < len(old); i++ {
		if new[i] != old[i] {
			res = append(res, i+1)
		}
	}
	return res
}

func AddTicket(ticket *model.Ticket) error {
	sql := "insert into ticket(userid,moveisessionid,state) values(?,?,1)"
	_, err := utils.Db.Exec(sql, ticket.UserId, ticket.MovieSessionId)
	if err != nil {
		return err
	}
	return nil
}

func GetTicketsByName(username string) ([]model.Ticket, error) {
	sql := "select user.username,cinema.cinemaname,moviesession.screenroom,ticket.seat,moviesession.showtime,movie.moviename,moviesession.price from ticket join moviesession on moviesession.moviesessionid=ticket.moviesessionid join user on user.userid=ticket.userid where user.username=?"
	rows, err := utils.Db.Query(sql, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tickets := []model.Ticket{}
	for rows.Next() {
		ticket := model.Ticket{}
		rows.Scan(&ticket.Owner, &ticket.Cinema, &ticket.Screen, &ticket.Seat, &ticket.Time, &ticket.Movie, &ticket.Price)
		tickets = append(tickets, ticket)
	}
	//fmt.Println("GetTicketsByName", tickets)
	return tickets, nil
}

func DeleteTicketByAllInfo(ticket *model.Ticket) {
	sql := "update ticket set state=0 where userid = ? and moviesessionid=? and seat=?"
	_, _ = utils.Db.Exec(sql, ticket.UserId, ticket.MovieSessionId, ticket.Seat)
}

func DeleteTicketWithoutSeat(ticket *model.Ticket) {
	sql := "update ticket set state=0 where userid = ? and moviesessionid=? and seat=?"
	_, _ = utils.Db.Exec(sql, ticket.UserId, ticket.MovieSessionId, ticket.Seat)
}

func GetMovieSessionByTicket(ticket *model.Ticket) (*model.MovieSession, error) {
	//sql := "select * from moviesession where cinema=? and screenroom=? and movietime=? and movie=? and price=?"
	sql0 := "select cinema.cinemaname,moviesession.screenroom,moviesession.showtime,movie.moviename,moviesession.showinfo,moviesession.price from moviesession join ticket on moviesession.moviesessionid=ticket.moviesessionid join movie on movie.movieid=moviesession.movieid where ticket.userid=?"
	//fmt.Println("GetMovieSessionByTicket", ticket)
	row := utils.Db.QueryRow(sql0, ticket.UserId)
	res := &model.MovieSession{}
	err := row.Scan(&res.ShowCinema, &res.ShowScreen, &res.ShowTime, &res.ShowMovie, &res.ShowInfo, &res.Price)
	//fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ModifySessionInfo(sessioninfo string, seat string, way string) string {
	pseats := strings.Split(strings.Replace(seat, "_", " ", -1), " ")
	runesessioninfo := []rune(sessioninfo)
	for _, pseat := range pseats {
		seatnum, err := strconv.Atoi(pseat)
		if err != nil || seatnum < 1 || seatnum > len(runesessioninfo) {
			continue
		}

		if way == "buy" {
			if runesessioninfo[seatnum-1] == '0' {
				runesessioninfo[seatnum-1] = '1'
			}
		} else if way == "sell" {
			if runesessioninfo[seatnum-1] == '1' {
				runesessioninfo[seatnum-1] = '0'
			}
		}
	}
	return string(runesessioninfo)
}

func GetAllTickets() []model.Ticket {
	//sql := "select * from tickets"
	sql0 := "select user.username,cinema.cinemaname,moviesession.screenroom,ticket.seat,moviesession.showtime,movie.moviename,moviesession.price from ticket join moviesession on ticket.moviesessionid=moviesession.moviesessionid join ticket.userid=user.userid where user.username=?"
	rows, _ := utils.Db.Query(sql0)
	defer rows.Close()
	tickets := []model.Ticket{}
	for rows.Next() {
		ticket := model.Ticket{}
		rows.Scan(&ticket.Owner, &ticket.Cinema, &ticket.Screen, &ticket.Seat, &ticket.Time, &ticket.Movie, &ticket.Price)
		tickets = append(tickets, ticket)
	}
	return tickets
}
