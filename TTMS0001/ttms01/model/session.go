package model

type Session struct {
	SessionID string
	Session   string
	UserName  string
	UserID    int
	IsLogin   bool
	State     bool
	Movies    []*Movie
	Cinema    []*Cinema
}
