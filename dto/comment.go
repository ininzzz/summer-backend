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

// blog/comment/post
type Blog_Comment_Post_ReqDTO struct {
	BlogID int64  `json:"blog_id"`
	Text   string `json:"text"`
	UserID int64
}

type Blog_Comment_Post_RespDTO struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
