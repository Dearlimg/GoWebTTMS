package dao

import (
	"ttms01/model"
	"ttms01/utils"
)

func AddComment(Comment *model.Comment) {
	sql := "insert into comment(movieid,userid,word,time,arid) values (?,?,?,?,?)"
	utils.Db.Exec(sql, Comment.MovieId, Comment.UserId, Comment.Word, Comment.Time, Comment.AtId, Comment.State)
}

func GetCommentsByMovieName(moviename string) []*model.Comment {
	sql := "select * from comment where movie=?"
	comments := []*model.Comment{}
	rows, _ := utils.Db.Query(sql, moviename)
	for rows.Next() {
		comment := &model.Comment{}
		rows.Scan(&comment.Movie, &comment.Speaker, &comment.Time, &comment.Word, &comment.At)
		//fmt.Println(comment)
		comments = append(comments, comment)
	}
	return comments
}

//func GetCommentsByMovieName(moviename string) []*model.Comment {
//	sql := "select * from comment where movieid=?"
//	Comments := []*model.Comment{}
//	rows, _ := utils.Db.Query(sql, moviename)
//	for rows.Next() {
//		Comment := &model.Comment{}
//		rows.Scan(sql, Comment.MovieId, Comment.UserId, Comment.Word, Comment.Time, Comment.AtId, Comment.State)
//		//fmt.Println(comment)
//		Comments = append(Comments, Comment)
//	}
//	return Comments
//}

func DeleteComment(Comment *model.Comment) {
	sql := "update comment set state=0 where id=?"
	utils.Db.Exec(sql, Comment.CommentId)
}
