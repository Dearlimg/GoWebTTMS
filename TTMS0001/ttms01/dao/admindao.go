package dao

import (
	"ttms01/model"
	"ttms01/utils"
)

func CheckAdmin(username string, password string) (*model.User, error) {
	sql := "select * from admin where adminname=? and password=?"
	user := &model.User{}
	row := utils.Db.QueryRow(sql, username, password)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func IsAdmin(username string) bool {
	sql := "select * from admin where username=?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
	fakestate := ""
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &fakestate)
	if user.ID > 0 {
		return true
	} else {
		return false
	}
}
