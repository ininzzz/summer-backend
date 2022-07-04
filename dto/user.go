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
	UserID string
}

type InfoResponseDTO struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Icon     string `json:"icon"`
}
