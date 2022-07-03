package model

type User struct {
	ID       int64
	Username string
	Password string
	Icon     []byte
	Gender   string
	Email    string
}
