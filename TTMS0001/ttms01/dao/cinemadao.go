package dao

import (
	"log"
	"ttms01/model"
	"ttms01/utils"
)

//func GetAllCinema() (*model.Page, error) {
//	sql1 := "select count(*) from cinema"
//	var totalRecord int64
//	row := utils.Db.QueryRow(sql1)
//	row.Scan(&totalRecord)
//
//	//fmt.Println(totalRecord)
//	sql := "select * from cinema"
//	var cinemas []*model.Cinema
//	rows, err := utils.Db.Query(sql)
//	if err != nil {
//		return nil, err
//	}
//	for rows.Next() {
//		cinema := &model.Cinema{}
//		rows.Scan(&cinema.CinemaId, &cinema.CinemaName, &cinema.Place, &cinema.Rank, &cinema.State)
//		cinemas = append(cinemas, cinema)
//	}
//	page := &model.Page{
//		Cinema:      cinemas,
//		TotalRecord: totalRecord,
//	}
//	//fmt.Println(cinemas[0].CinemaName, "MEYOUMNA")
//	return page, nil
//}

//func GetAllCinema() (*model.Page, error) {
//	sql1 := "select count(*) from cinema"
//	var totalRecord int64
//	row := utils.Db.QueryRow(sql1)
//	row.Scan(&totalRecord)
//
//	//fmt.Println(totalRecord)
//	sql := "select cinema(cinemaname,place,cinemarank) from cinema where state=1"
//	var cinemas []*model.Cinema
//	rows, err := utils.Db.Query(sql)
//	if err != nil {
//		return nil, err
//	}
//	for rows.Next() {
//		cinema := &model.Cinema{}
//		rows.Scan(&cinema.CinemaName, &cinema.Place, &cinema.Rank)
//		fmt.Println(cinema)
//		cinemas = append(cinemas, cinema)
//	}
//	page := &model.Page{
//		Cinema:      cinemas,
//		TotalRecord: totalRecord,
//	}
//	//fmt.Println(cinemas[0].CinemaName, "MEYOUMNA")
//	return page, nil
//}

func GetAllCinema() (*model.Page, error) {
	sql1 := "SELECT COUNT(*) FROM cinema"
	var totalRecord int64
	row := utils.Db.QueryRow(sql1)
	if err := row.Scan(&totalRecord); err != nil {
		return nil, err
	}

	sql := "SELECT cinemaname, place, cinemarank FROM cinema WHERE state = 1"
	var cinemas []*model.Cinema
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 确保 rows 资源在函数结束时被关闭

	for rows.Next() {
		cinema := &model.Cinema{}
		if err := rows.Scan(&cinema.CinemaName, &cinema.Place, &cinema.Rank); err != nil {
			return nil, err
		}
		cinemas = append(cinemas, cinema)
	}

	// 检查 rows.Err() 以确保没有错误
	if err := rows.Err(); err != nil {
		return nil, err
	}

	page := &model.Page{
		Cinema:      cinemas,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetCinemaByCondition(area string, rank string) []*model.Cinema {
	sql := "SELECT cinemaname, place, cinemarank FROM cinema WHERE state = 1"
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

//func GetCinemaByCondition(area string, rank string) []*model.Cinema {
//	sql := "select * from cinema where 1 = 1"
//	var args []interface{}
//
//	if area != "" {
//		sql += " and place = ?"
//		args = append(args, area)
//	}
//	if rank != "" {
//		sql += " and cinemarank = ?"
//		args = append(args, rank)
//	}
//
//	//fmt.Println("Executing query:", sql)
//
//	rows, err := utils.Db.Query(sql, args...)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	var cinemas []*model.Cinema
//	for rows.Next() {
//		cinema := &model.Cinema{}
//		if err := rows.Scan(&cinema.CinemaId, &cinema.CinemaName, &cinema.Place, &cinema.Rank, &cinema.State); err != nil {
//			log.Println("Scan error:", err)
//			continue
//		}
//		cinemas = append(cinemas, cinema)
//	}
//	if err := rows.Err(); err != nil {
//		log.Println("Rows error:", err)
//	}
//
//	return cinemas
//}

func AddCinema(cinema *model.Cinema) {
	sql := "insert into cinema(cinemaname, place, cinemarank,state) values(?, ?, ?,1)"
	utils.Db.Exec(sql, cinema.CinemaName, cinema.Place, cinema.Rank)
}

func DeleteCinemaByCinemaName(cinemaName string) {
	//sql := "delete from cinema where cinema_name = ?"
	//utils.Db.Exec(sql, cinemaName)
	sql := "update cinema set state=0 where cinemaname=?"
	utils.Db.Exec(sql, cinemaName)
}

func SaveCinema(cinema *model.Cinema, cinemaName string) {
	sql := "UPDATE cinema SET cinemarank=?, place=?, cinemaname=? WHERE cinemaname=?"
	utils.Db.Exec(sql, cinema.Rank, cinema.Place, cinema.CinemaName, cinemaName)
}

func GetAllPlcae() []*string {
	sql := "select distinct place from cinema"
	rows, _ := utils.Db.Query(sql)
	var genres []*string
	for rows.Next() {
		var genre *string
		rows.Scan(&genre)
		genres = append(genres, genre)
	}
	return genres
}

func GetAllRank() []*string {
	sql := "select distinct cinema_rank from cinema"
	rows, _ := utils.Db.Query(sql)
	var genres []*string
	for rows.Next() {
		var genre *string
		rows.Scan(&genre)
		genres = append(genres, genre)
	}
	return genres
}
