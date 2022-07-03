package model

import "github.com/ininzzz/summer-backend/dto"

type Blog struct {
	ID      int64
	UserID  int64
	Title   string
	Text    string
	Like    int
	Comment []dto.BlogCommentListResponseDTO
}
