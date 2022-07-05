package model

type Comment struct {
	CommentID       int64
	BlogID          int64
	UserID          int64
	Text            string
	CreateTimeStamp int64
	ModifyTimeStamp int64
}
