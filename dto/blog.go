package dto


// blog/info
type BlogInfoRequestDTO struct {
	BlogID int64
}

type BlogInfoResponseDTO struct {
	UserID     int64                        `json:"user_id"`
	UserName   string                       `json:"user_name"`
	UserAvatar string                       `json:"user_avatar"`
	Text       string                       `json:"text"`
	Imgs       string                       `json:"imgs"`
	Like       int64                        `json:"like"`
	Comments   []BlogCommentListResponseDTO `json:"comments"`
}

// blog/home/list
type BlogHomeListRequestDTO struct {
	LastTimeStamp int64
}

type HomeListBlog struct {
	BlogID     int64  `json:"blog_id"`
	UserID     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	Text       string `json:"text"`
	Imgs       string `json:"imgs"`
	Like       int64  `json:"like"`
}
type BlogHomeListResponseDTO struct {
	LastTimeStamp int64          `json:"lastTimeStamp"`
	BlogList      []HomeListBlog `json:"blog_list"`
}

// blog/space/list
type BlogSpaceListRequestDTO struct {
	UserID int64
}

type BlogSpaceListResponseDTO struct {
	BlogID     int64  `json:"blog_id"`
	UserID     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	Text       string `json:"text"`
	Imgs       string `json:"imgs"`
	Like       int64  `json:"like"`