package dao

import (
	"strconv"
	"strings"
	"ttms01/model"
	"ttms01/utils"
)

func SaveData(data string, session *model.MovieSession) error {
	sql := "update movie_session set info=? where cinema=? and screenroom=? and movietime=? and movie=?"
	_, err := utils.Db.Exec(sql, data, session.ShowCinema, session.ShowScreen, session.ShowTime, session.ShowMovie)
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
	sql := "insert into tickets(owner,cinema,screen,seat,showtime,movie,price) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, ticket.Owner, ticket.Cinema, ticket.Screen, ticket.Seat, ticket.Time, ticket.Movie, ticket.Price)
	if err != nil {
		return err
	}
	return nil
}

func GetTicketsByName(username string) ([]model.Ticket, error) {
	sql := "select * from tickets where owner=?"
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

func DeleteTicketByAllInfo(ticket *model.Ticket) error {
	sql := "delete from tickets where owner=? and cinema=? and screen=? and seat=? and showtime=? and movie=? and price=?"
	//fmt.Println("DeleteTicketByAllInfo", ticket)
	_, err := utils.Db.Exec(sql, ticket.Owner, ticket.Cinema, ticket.Screen, ticket.Seat, ticket.Time, ticket.Movie, ticket.Price)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTicketWithoutSeat(ticket *model.Ticket) error {
	sql := "delete from tickets where owner=? and cinema=? and screen=? and showtime=? and movie=? and price=?"
	//fmt.Println("DeleteTicketByAllInfo", ticket)
	_, err := utils.Db.Exec(sql, ticket.Owner, ticket.Cinema, ticket.Screen, ticket.Time, ticket.Movie, ticket.Price)
	if err != nil {
		return err
	}
	return nil
}

func GetMovieSessionByTicket(ticket *model.Ticket) (*model.MovieSession, error) {
	sql := "select * from movie_session where cinema=? and screenroom=? and movietime=? and movie=? and price=?"
	//fmt.Println("GetMovieSessionByTicket", ticket)
	row := utils.Db.QueryRow(sql, ticket.Cinema, ticket.Screen, ticket.Time, ticket.Movie, ticket.Price)
	res := &model.MovieSession{}
	err := row.Scan(&res.ShowCinema, &res.ShowScreen, &res.ShowTime, &res.ShowMovie, &res.ShowInfo, &res.Price)
	//fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//func ModifySessionInfo(sessioninfo string, seat string) string {
//	pseat := strings.Replace(seat, "_", "", -1)
//	seatnums := []int(pseat)
//	runesessioninfo := []int(sessioninfo)
//	fmt.Println(seatnums)
//	fmt.Println(runesessioninfo)
//	for i := 0; i < len(seatnums); i++ {
//		for j := 0; j < len(sessioninfo); j++ {
//			if j == seatnums[i]-1 {
//				if runesessioninfo[j] == 0 {
//					runesessioninfo[j] = 1
//				} else {
//					runesessioninfo[j] = 0
//				}
//			}
//		}
//	}
//}

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
	sql := "select * from tickets"
	rows, _ := utils.Db.Query(sql)
	defer rows.Close()
	tickets := []model.Ticket{}
	for rows.Next() {
		ticket := model.Ticket{}
		rows.Scan(&ticket.Owner, &ticket.Cinema, &ticket.Screen, &ticket.Seat, &ticket.Time, &ticket.Movie, &ticket.Price)
		tickets = append(tickets, ticket)
	}
	return tickets
}
