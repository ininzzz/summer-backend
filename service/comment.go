package service

import (
	"context"

	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/sirupsen/logrus"
)

var CommentService commentService

type commentService struct {
	commentRepo infra.CommentRepo
}

// blog/comment/list
func (u *commentService) CommentList(ctx context.Context, reqDTO *dto.BlogCommentListRequestDTO) (*common.Response, error) {
	comments, err := u.commentRepo.Find(ctx, &infra.CommentQuery{
		BlogID: &reqDTO.BlogID,
	})
	if err != nil {
		logrus.Errorf("[commentService CommentList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := []*dto.BlogCommentListResponseDTO{}
	for _, comment := range comments {
		user, _ := UserService.userRepo.FindByID(ctx, &infra.UserQuery{
			ID: &comment.UserID,
		})
		data = append(data, &dto.BlogCommentListResponseDTO{
			Text:       comment.Text,
			UserID:     comment.UserID,
			UserName:   user.Username,
			UserAvatar: user.UserAvatar,
		})
	}
	return common.NewResponseOfSuccess(data), nil
}
