package dto

// user/login
type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// user/info
type InfoRequestDTO struct {
	UserID int64
}

type InfoResponseDTO struct {
	Username string `json:"user_name"`
	Gender   string `json:"user_gender"`
	Email    string `json:"user_email"`
	Avatar   string `json:"user_avatar"`
}

// user/email/code
type User_Email_Code_ReqDTO struct {
	Email string `json:"email"`
}

type User_Email_Code_RespDTO struct {
	Ok bool `json:"ok"`
}
