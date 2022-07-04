package dto

// blog/comment/list
type BlogCommentListRequestDTO struct {
	BlogID int64
}

type BlogCommentListResponseDTO struct {
	UserID     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	Text       string `json:"text"`
}
