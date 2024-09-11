package model

type Comment struct {
	CommentId  string
	Movie      string
	MovieId    string
	UserId     string
	Speaker    string
	Word       string
	Time       string
	AtId       string
	At         string
	State      string
	SonComment []*SonComment
}

type SonComment struct {
	SonCommentId    string
	FatherCommentId string
	Replier         string
	Replies         string
	Time            string
}
