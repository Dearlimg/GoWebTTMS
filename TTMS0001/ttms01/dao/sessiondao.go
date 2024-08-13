package dao

import (
	"net/http"
	"ttms01/model"
	"ttms01/utils"
)

func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values (?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id=?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessionID string) (*model.Session, error) {
	sqlStr := "select * from sessions where session_id=?"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := inStmt.QueryRow(sessionID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.UserName, &session.UserID)
	return session, nil
}

func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	return false, nil
}
