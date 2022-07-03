package dto

// blog/all/list
type BlogListAllRequestDTO struct {
}

type BlogListAllResponseDTO struct {
	ID     int64 `json:"id"`
	UserID int64
	Title  string `json:"title"`
	Like   int    `json:"like"`
}

// blog/list
type BlogListRequestDTO struct {
	UserID string `query:"user_id"`
}

type BlogListResponseDTO struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Like  int    `json:"like"`
}

// blog/info
type BlogInfoRequestDTO struct {
	BlogID string `param:"blog_id"`
}

type BlogInfoResponseDTO struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Like  int    `json:"like"`
}

// blog/post
type BlogPostRequestDTO struct {
	UserID string
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type BlogPostResponseDTO struct {
}

// blog/comment/post
type BlogCommentPostRequestDTO struct {
	UserID  string
	BlogID  int64  `json:"blog_id"`
	Comment string `json:"comment"`
}

type BlogCommentPostResponseDTO struct {
}

// blog/comment/list
type BlogCommentListRequestDTO struct {
	BlogID string `query:"blog_id"`
}

type BlogCommentListResponseDTO struct {
	Username string `json:"username"`
	Comment  string `json:"comment"`
}

// blog/like
type BlogLikeRequestDTO struct {
	BlogID int64 `json:"blog_id"`
	Value  int   `json:"value"`
}

type BlogLikeResponseDTO struct {
}
