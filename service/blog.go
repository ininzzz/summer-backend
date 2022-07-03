package service

import (
	"context"
	"strconv"

	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/sirupsen/logrus"
)

var BlogService blogService

type blogService struct {
	blogRepo infra.BlogRepo
}

func (u *blogService) List(ctx context.Context, blogListDTO *dto.BlogListRequestDTO) (*common.Response, error) {
	blogs, err := u.blogRepo.Find(ctx, &infra.BlogQuery{
		UserID: &blogListDTO.UserID,
	})
	if err != nil {
		logrus.Errorf("[blogService BlogList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := []*dto.BlogListResponseDTO{}
	for _, blog := range blogs {
		data = append(data, &dto.BlogListResponseDTO{
			ID:    blog.ID,
			Title: blog.Title,
		})
	}
	return common.NewResponseOfSuccess(data), nil
}

func (u *blogService) Info(ctx context.Context, blogInfoDTO *dto.BlogInfoRequestDTO) (*common.Response, error) {
	blogID, err := strconv.Atoi(blogInfoDTO.BlogID)
	blogID64 := int64(blogID)
	if err != nil {
		logrus.Errorf("[blogService BlogInfo] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	blogs, err := u.blogRepo.Find(ctx, &infra.BlogQuery{
		ID: &blogID64,
	})
	if err != nil {
		logrus.Errorf("[blogService BlogInfo] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.BlogInfoResponseDTO{
		Title: blogs[0].Title,
		Text:  blogs[0].Text,
	}
	return common.NewResponseOfSuccess(data), nil
}
