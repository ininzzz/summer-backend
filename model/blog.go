package model

type Blog struct {
	BlogID          int64
	UserID          int64
	Text            string
	Imgs            string
	Like            int
	CreateTimestamp int64
}
