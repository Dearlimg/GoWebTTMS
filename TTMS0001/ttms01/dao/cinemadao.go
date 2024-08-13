package dao

import (
	"log"
	"ttms01/model"
	"ttms01/utils"
)

func GetAllCinema() (*model.Page, error) {
	sql1 := "select count(*) from cinema"
	var totalRecord int64
	row := utils.Db.QueryRow(sql1)
	row.Scan(&totalRecord)

	//fmt.Println(totalRecord)
	sql := "select * from cinema"
	var cinemas []*model.Cinema
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cinema := &model.Cinema{}
		rows.Scan(&cinema.CinemaName, &cinema.Place, &cinema.Rank)
		cinemas = append(cinemas, cinema)
	}
	page := &model.Page{
		Cinema:      cinemas,
		TotalRecord: totalRecord,
	}
	//fmt.Println(cinemas[0].CinemaName, "MEYOUMNA")
	return page, nil
}

func GetCinemaByCondition(area string, rank string) []*model.Cinema {
	sql := "select * from cinema where 1 = 1"
	var args []interface{}

	if area != "" {
		sql += " and place = ?"
		args = append(args, area)
	}
	if rank != "" {
		sql += " and cinema_rank = ?"
		args = append(args, rank)
	}

	//fmt.Println("Executing query:", sql)

	rows, err := utils.Db.Query(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cinemas []*model.Cinema
	for rows.Next() {
		cinema := &model.Cinema{}
		if err := rows.Scan(&cinema.CinemaName, &cinema.Place, &cinema.Rank); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		cinemas = append(cinemas, cinema)
	}
	if err := rows.Err(); err != nil {
		log.Println("Rows error:", err)
	}

	return cinemas
}

func AddCinema(cinema *model.Cinema) {
	sql := "insert into cinema(cinema_name, place, cinema_rank) values(?, ?, ?)"
	utils.Db.Exec(sql, cinema.CinemaName, cinema.Place, cinema.Rank)
}

func DeleteCinemaByCinemaName(cinemaName string) {
	sql := "delete from cinema where cinema_name = ?"
	utils.Db.Exec(sql, cinemaName)
}

func SaveCinema(cinema *model.Cinema, cinemaName string) {
	sql := "UPDATE cinema SET cinema_rank=?, place=?, cinema_name=? WHERE cinema_name=?"
	//fmt.Println(cinema.Rank)
	utils.Db.Exec(sql, cinema.Rank, cinema.Place, cinema.CinemaName, cinemaName)
}
