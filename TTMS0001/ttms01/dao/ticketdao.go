package dao

import (
	"fmt"
	"strconv"
	"strings"
	"ttms01/model"
	"ttms01/utils"
)

func SaveData(data string, session *model.MovieSession) error {
	sql := "update moviesession set showinfo=? where screenroom=? and showtime=?"
	//sql := "update moviesession set showinfo=? where cinemaid=? and screenroom=? and showtime=? and movieid=?"
	fmt.Println("SaveData", data, session.CinemaId, session.ShowScreen, session.ShowTime, session.MovieId)
	_, err := utils.Db.Exec(sql, data, session.ShowScreen, session.ShowTime)
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

//func AddTicket(ticket *model.Ticket) error {
//	sql0 := "select user.userid,moviesession.moviesessionid from ticket join user on user.userid=ticket.userid join moviesession on moviesession.moviesessionid=ticket.moviesessionid where user.username=?"
//	utils.Db.Query(sql0, ticket.Owner)
//	fmt.Println(ticket.UserId, ticket.MovieSessionId, ticket.Seat)
//	sql := "insert into ticket(userid,moviesessionid,state,seat) values(?,?,1,?)"
//	_, err := utils.Db.Exec(sql, ticket.UserId, ticket.MovieSessionId, ticket.Seat)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func AddTicket(ticket *model.Ticket) error {
//	// 定义用于查询的 SQL 语句
//	//sql0 := `
//	//	SELECT user.userid, moviesession.moviesessionid
//	//	FROM ticket
//	//	JOIN user ON user.userid = ticket.userid
//	//	JOIN moviesession ON moviesession.moviesessionid = ticket.moviesessionid
//	//	WHERE user.username = ?`
//
//	//// 执行查询，并将结果存储在变量中
//	//row := utils.Db.QueryRow(sql0, ticket.Owner)
//
//	// 定义用于接收查询结果的变量
//	//var userId int
//	//var movieSessionId int
//	//
//	//// 执行扫描，将结果赋值给变量
//	//err := row.Scan(&userId, &movieSessionId)
//	//fmt.Println("AddTicket", userId, movieSessionId)
//	//if err != nil {
//	//	// 检查错误是否是因为未找到结果
//	//	if err == sql.ErrNoRows {
//	//		return fmt.Errorf("no such user or movie session found for username: %s", ticket.Owner)
//	//	}
//	//	return err
//	//}
//
//	// 打印调试信息
//	//fmt.Println(userId, movieSessionId, ticket.Seat)
//	//
//	//// 定义插入的 SQL 语句
//	//sql := "INSERT INTO ticket(userid, moviesessionid, state, seat) VALUES (?, ?, 1, ?)"
//	//
//	//// 执行插入操作
//	//_, err = utils.Db.Exec(sql, userId, movieSessionId, ticket.Seat)
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//return nil
//}

func AddTicket(ticket *model.Ticket) {
	fmt.Println("			AddTicket1", ticket)
	var id int
	var movieSessionID int

	if ticket.Owner == "admin" {
		sql0 := "SELECT adminid FROM admin WHERE adminname=?"
		err := utils.Db.QueryRow(sql0, ticket.Owner).Scan(&id)
		fmt.Println("wad 	d	dwa		dwa", id)
		if err != nil {
			fmt.Println("Error fetching admin ID:", err)
			return
		}

		sql3 := "SELECT moviesessionid FROM moviesession WHERE showtime=? AND price=?"
		err = utils.Db.QueryRow(sql3, ticket.Time, ticket.Price).Scan(&movieSessionID)
		if err != nil {
			fmt.Println("Error fetching movie session ID:", err)
			return
		}

		sql2 := "INSERT INTO ticket(userid, moviesessionid, state, seat, owner) VALUES(?,?,?,?,?)"
		fmt.Println("AddTicket", id, movieSessionID, 1, ticket.Seat, ticket.Owner)
		_, err = utils.Db.Exec(sql2, id, movieSessionID, 1, ticket.Seat, ticket.Owner)
		fmt.Println(" 			AddTicket2", id, movieSessionID, 1, ticket.Seat, ticket.Owner)
		if err != nil {
			fmt.Println("Error inserting ticket:", err)
			return
		}

	} else {
		sql1 := "SELECT userid FROM user WHERE username=?"
		err := utils.Db.QueryRow(sql1, ticket.Owner).Scan(&id)
		if err != nil {
			fmt.Println("Error fetching user ID:", err)
			return
		}

		sql3 := "SELECT moviesessionid FROM moviesession WHERE showtime=? AND price=?"
		err = utils.Db.QueryRow(sql3, ticket.Time, ticket.Price).Scan(&movieSessionID)
		if err != nil {
			fmt.Println("Error fetching movie session ID:", err)
			return
		}

		sql2 := "INSERT INTO ticket(userid, moviesessionid, state, seat, owner) VALUES(?,?,?,?,?)"
		_, err = utils.Db.Exec(sql2, id, movieSessionID, 1, ticket.Seat, ticket.Owner)
		if err != nil {
			fmt.Println("Error inserting ticket:", err)
			return
		}
	}
}

//func AddTicket(ticket *model.Ticket) {
//	fmt.Println("AddTicket", ticket)
//	if ticket.Owner == "admin" {
//		sql0 := "select adminid from admin where adminname=?"
//		utils.Db.QueryRow(sql0, ticket.Movie)
//	} else {
//		sql1 := "select userid from user where username=?"
//		utils.Db.QueryRow(sql1)
//
//		sql3 := "select moviesessionid from moviesession where showtime=? price=?"
//		utils.Db.QueryRow(sql3, ticket.Time, ticket.Price)
//
//		sql2 := "insert into ticket(userid,moviesessionid,state,seat,owner) values(?,?,1,?,?)"
//		utils.Db.Exec(sql2)
//
//	}
//}

func GetTicketsByName(username string) ([]model.Ticket, error) {
	//sql := "select user.username,cinema.cinemaname,moviesession.screenroom,ticket.seat,moviesession.showtime,movie.moviename,moviesession.price from ticket join moviesession on moviesession.moviesessionid=ticket.moviesessionid join user on user.userid=ticket.userid where user.username=?"
	sql := "select user.username,cinema.cinemaname,moviesession.screenroom,ticket.seat,moviesession.showtime,movie.moviename,moviesession.price from ticket join moviesession on ticket.moviesessionid=moviesession.moviesessionid join user on ticket.userid=user.userid join cinema on cinema.cinemaid=moviesession.cinemaid join movie on moviesession.movieid = movie.movieid where user.username=? and ticket.state>0"
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

//func DeleteTicketByAllInfo(ticket *model.Ticket) {
//	sql0 := "select moviesessionid from moviesession where showtime=? and screenroom=? and price=?"
//	utils.Db.Exec(sql0, ticket.Time, ticket.Screen, ticket.Price)
//
//	fmt.Println("Delet eTicke  tByAllInfo", ticket.UserId, ticket.MovieSessionId, ticket.Seat)
//	sql := "update ticket set state=0 where moviesessionid=? and seat=?"
//	_, _ = utils.Db.Exec(sql, ticket.MovieSessionId, ticket.Seat)
//}

func DeleteTicketByAllInfo(ticket *model.Ticket) {
	// 定义 SQL 查询语句，获取 moviesessionid
	sql0 := "SELECT moviesessionid FROM moviesession WHERE showtime=? AND screenroom=? AND price=?"

	// 定义一个变量用来保存查询结果
	var movieSessionID int

	// 执行查询并将结果赋值给 movieSessionID
	err := utils.Db.QueryRow(sql0, ticket.Time, ticket.Screen, ticket.Price).Scan(&movieSessionID)
	if err != nil {
		// 处理查询错误，例如记录日志或返回错误
		fmt.Println("Error querying moviesession:", err)
		return
	}

	// 打印日志信息，帮助调试
	fmt.Println("DeleteTicketByAllInfo", movieSessionID, ticket.Seat)

	// 定义 SQL 更新语句，将票务状态更新为 0
	sql := "UPDATE ticket SET state=0 WHERE moviesessionid=? AND seat=?"

	// 执行更新操作
	_, err = utils.Db.Exec(sql, movieSessionID, ticket.Seat)
	if err != nil {
		// 处理更新错误，例如记录日志或返回错误
		fmt.Println("Error updating ticket:", err)
		return
	}

	//

}

func DeleteTicketWithoutSeat(ticket *model.Ticket) {
	sql := "update ticket set state=0 where userid = ? and moviesessionid=? and seat=?"
	_, _ = utils.Db.Exec(sql, ticket.UserId, ticket.MovieSessionId, ticket.Seat)
}

//func GetMovieSessionByTicket(ticket *model.Ticket) (*model.MovieSession, error) {
//	//sql := "select * from moviesession where cinema=? and screenroom=? and movietime=? and movie=? and price=?"
//	//fmt.Println("GetMovieSessionByTicket", ticket.UserId)
//	sql01 := "select userid from user where username=?"
//	var userid string
//	err0 := utils.Db.QueryRow(sql01, ticket.Owner).Scan(&userid)
//	if err0 != nil {
//		panic(err0)
//	}
//	fmt.Println("GetMovieSessionByTicket", userid, "wad", ticket)
//
//	sql0 := "select cinema.cinemaname,moviesession.screenroom,moviesession.showtime,movie.moviename,moviesession.showinfo,moviesession.price from moviesession join ticket on moviesession.moviesessionid=ticket.moviesessionid join movie on movie.movieid=moviesession.movieid join cinema on cinema.cinemaid=moviesession.moviesessionid where moviesession.showtime=? and moviesession.price=? and moviesession.screenroom=?"
//
//	//fmt.Println("GetMovieSessionByTicket", ticket)
//	row := utils.Db.QueryRow(sql0, ticket.Time, ticket.Price, ticket.Screen)
//	res := &model.MovieSession{}
//	err := row.Scan(&res.ShowCinema, &res.ShowScreen, &res.ShowTime, &res.ShowMovie, &res.ShowInfo, &res.Price)
//	fmt.Println("GetMovieS ess ionB yTicket", res)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

//func GetMovieSessionByTicket(ticket *model.Ticket) (*model.MovieSession, error) {
//	sql := "select moviesessionid from ticket where ticketid=?"
//	var movieSessionID int
//	err := utils.Db.QueryRow(sql, ticket.TicketId).Scan(&movieSessionID)
//	if err != nil {
//		panic(err)
//	}
//	var res *model.MovieSession
//	sql1 := "select cinema.cinemaname,moviesession.screenroom,moviesession.showtime,movie.moviename,moviesession.showinfo,moviesession.price from moviesession join cinema on moviesession.cinemaid=cinema.cinemaid join cinema on cinema.cinemaid=moviesession.cinemaid where moviesession.moviesessionid=?"
//	utils.Db.QueryRow(sql1, movieSessionID).Scan(&res.ShowCinema, &res.ShowScreen, &res.ShowTime, &res.ShowMovie, &res.ShowInfo, &res.Price)
//	fmt.Println("res", res)
//	return res, nil
//}

func GetMovieSessionByTicket(ticket *model.Ticket) (*model.MovieSession, error) {
	sql := "SELECT moviesessionid FROM ticket WHERE ticketid=?"
	var movieSessionID int
	err := utils.Db.QueryRow(sql, ticket.TicketId).Scan(&movieSessionID)
	if err != nil {
		return nil, err
	}

	fmt.Println(movieSessionID)
	var res model.MovieSession
	sql1 := `
		SELECT cinema.cinemaname, moviesession.screenroom, moviesession.showtime, movie.moviename, moviesession.showinfo, moviesession.price
		FROM moviesession
		JOIN cinema ON moviesession.cinemaid = cinema.cinemaid
		JOIN movie ON movie.movieid = moviesession.movieid
		WHERE moviesession.moviesessionid = ?`
	err = utils.Db.QueryRow(sql1, movieSessionID).Scan(&res.ShowCinema, &res.ShowScreen, &res.ShowTime, &res.ShowMovie, &res.ShowInfo, &res.Price)
	if err != nil {
		return nil, err
	}

	fmt.Println("rdwawadwadawdawes", res)
	return &res, nil
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
	fmt.Println("GetAllTickets1				")
	sql0 := "select ticket.ticketid,user.username,cinema.cinemaname,moviesession.screenroom,ticket.seat,moviesession.showtime,movie.moviename,moviesession.price from ticket join moviesession on ticket.moviesessionid=moviesession.moviesessionid join user on ticket.userid=user.userid join cinema on cinema.cinemaid=moviesession.cinemaid join movie on moviesession.movieid = movie.movieid where ticket.state>0"
	rows, _ := utils.Db.Query(sql0)
	fmt.Println("GetAllTickets1		2		")
	defer rows.Close()
	tickets := []model.Ticket{}
	for rows.Next() {
		ticket := model.Ticket{}
		rows.Scan(&ticket.TicketId, &ticket.Owner, &ticket.Cinema, &ticket.Screen, &ticket.Seat, &ticket.Time, &ticket.Movie, &ticket.Price)
		//fmt.Println("GetAllTickets				", ticket)
		tickets = append(tickets, ticket)
	}
	return tickets
}
