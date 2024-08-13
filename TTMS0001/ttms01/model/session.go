package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int
	IsLogin   bool
	Movies    []*Movie
	Cinema    []*Cinema
}
