package dao

import (
	"crypto/md5"
	"encoding/hex"
	"ttms01/model"
	"ttms01/utils"
)

func AddUser(user model.User) error {
	sql := "insert into users(username,password,email,state) values(?,?,?,1)"

	md5Pwd := md5.Sum([]byte(user.Password))
	strmd5 := hex.EncodeToString(md5Pwd[:])

	_, err := utils.Db.Exec(sql, user.Username, strmd5, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserByUserName(username string) {
	sql0 := "select userid from user where username=?"
	rows := utils.Db.QueryRow(sql0, username)
	sql := "update user set state=0 where userid=?"
	utils.Db.Exec(sql, rows)
}

func ModifyUserPassWordByUserName(username string) error {
	sql := "update user set password=? where username=?"
	_, err := utils.Db.Exec(sql, username)
	if err != nil {
		return err
	}
	return nil
}

func SearchUserByUserName(username string) *model.User {
	sql := "select * from user where username=?"
	rows, err := utils.Db.Query(sql, username)
	if err != nil {
		return nil
	}
	for rows.Next() {
		var user *model.User
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if user.Username == username {
			return user
		}
	}
	return nil
}

func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select userid,username,password,email from user where username=?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlStr := "select userid,username,password,email from users where username=? and password=?"

	md5Pwd := md5.Sum([]byte(password))
	strmd5 := hex.EncodeToString(md5Pwd[:])

	row := utils.Db.QueryRow(sqlStr, username, strmd5)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	user.Password = password
	return user, nil
}
