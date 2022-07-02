package dto

// blog/list
type BlogListRequestDTO struct {
	UserID int64
}

type BlogListResponseDTO struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// blog/:blog_id
type BlogInfoRequestDTO struct {
	BlogID string `param:"blog_id"`
}

type BlogInfoResponseDTO struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
