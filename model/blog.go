package model

type Blog struct {
	BlogID          int64
	UserID          int64
	Text            string
	Imgs            string
	CreateTimestamp int64
	ModifyTimestamp int64
}