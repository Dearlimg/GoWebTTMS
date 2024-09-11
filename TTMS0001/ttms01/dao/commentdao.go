package dao

import (
	"fmt"
	"ttms01/model"
	"ttms01/utils"
)

func AddComment(Comment *model.Comment) error {
	// 使用查询获取用户ID
	sql0 := "select userid from user where username=?"
	var userid string
	err := utils.Db.QueryRow(sql0, Comment.Speaker).Scan(&userid)
	if err != nil {
		return fmt.Errorf("failed to get user ID: %w", err)
	}

	// 将获取到的userid赋值到Comment.UserId
	Comment.UserId = userid

	// 插入评论到数据库
	sql := "insert into comment(movieid, userid, word, time, state) values (?, ?, ?, ?, 1)"
	_, err = utils.Db.Exec(sql, Comment.MovieId, Comment.UserId, Comment.Word, Comment.Time)
	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}

	return nil
}

//func AddComment(Comment *model.Comment) {
//	sql0 := "select userid from user where username=?"
//	utils.Db.QueryRow(sql0, Comment.Speaker)
//	var userid string
//
//	sql := "insert into comment(movieid,userid,word,time,state) values (?,?,?,?,1)"
//	fmt.Println(Comment.MovieId, Comment.UserId, Comment.Word, Comment.Time)
//	utils.Db.Exec(sql, Comment.MovieId, Comment.UserId, Comment.Word, Comment.Time)
//}

func GetCommentsByMovieName(moviename string) []*model.Comment {
	//已经传成id了
	sql := "select comment.commentid,movie.moviename,user.username,comment.time,comment.word,user.username from comment join user on user.userid=comment.userid join movie on movie.movieid=comment.movieid where comment.movieid=?"
	comments := []*model.Comment{}
	rows, _ := utils.Db.Query(sql, moviename)
	for rows.Next() {
		comment := &model.Comment{}
		rows.Scan(&comment.CommentId, &comment.Movie, &comment.Speaker, &comment.Time, &comment.Word, &comment.At)
		comment.SonComment = GetAllSonComment(comment.CommentId)
		fmt.Println(comment.CommentId, comment.Speaker, comment.Time, comment.Word, comment.SonComment)
		comments = append(comments, comment)
	}
	return comments
}

func GetAllSonComment(FatherCommentId string) []*model.SonComment {
	sql := "select * from soncomment where fatherconnentid=?"
	var SonComments []*model.SonComment
	rows, _ := utils.Db.Query(sql, FatherCommentId)
	for rows.Next() {
		SonComment := &model.SonComment{}
		rows.Scan(&SonComment.SonCommentId, &SonComment.FatherCommentId, &SonComment.Replies, &SonComment.Time, &SonComment.Replier)
		SonComment.Replier = GetSonCommentName(SonComment.Replier)
		SonComments = append(SonComments, SonComment)
	}
	return SonComments
}

func GetSonCommentName(SonCommentId string) string {
	sql := "select username from user where userid=?"
	var username string
	rows := utils.Db.QueryRow(sql, SonCommentId)
	rows.Scan(&username)
	return username
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
