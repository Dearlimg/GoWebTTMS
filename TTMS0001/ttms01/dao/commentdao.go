package dao

import (
	"ttms01/model"
	"ttms01/utils"
)

func AddComment(Comment *model.Comment) {
	sql := "insert into comment(movie,speaker,time,word,at) values (?,?,?,?,?)"
	utils.Db.Exec(sql, Comment.Movie, Comment.Speaker, Comment.Time, Comment.Comment, Comment.At)
}

func GetCommentsByMovieName(moviename string) []*model.Comment {
	sql := "select * from comment where movie=?"
	comments := []*model.Comment{}
	rows, _ := utils.Db.Query(sql, moviename)
	for rows.Next() {
		comment := &model.Comment{}
		rows.Scan(&comment.Movie, &comment.Speaker, &comment.Time, &comment.Comment, &comment.At)
		//fmt.Println(comment)
		comments = append(comments, comment)
	}
	return comments
}

func DeleteComment(Comment *model.Comment) {
	sql := "delete from comment where movie=? and speaker=? and time=? and word=?"
	utils.Db.Exec(sql, Comment.Movie, Comment.Speaker, Comment.Time, Comment.Comment)
}
