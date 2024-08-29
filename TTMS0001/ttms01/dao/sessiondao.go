package dao

import (
	"net/http"
	"ttms01/model"
	"ttms01/utils"
)

func AddSession(sess *model.Session) error {
	sqlStr := "insert into session(userid,state,session) values (?,1,?)"
	_, err := utils.Db.Exec(sqlStr, sess.UserID, sess.SessionID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) {
	SQL := "update session set state=0 where session = ?"
	utils.Db.Exec(SQL, sessID)
}

func GetSession(sessionID string) (*model.Session, error) {
	sqlStr := "select * from session where session=?"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := inStmt.QueryRow(sessionID)
	session := &model.Session{}
	row.Scan(&session.SessionID, &session.UserID, &session.State, &session.Session)
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
