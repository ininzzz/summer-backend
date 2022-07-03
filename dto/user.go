package dto

// user/login
type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

// user/info
type InfoRequestDTO struct {
	UserID int64
}

type InfoResponseDTO struct {
	Username string `json:"username"`
}


