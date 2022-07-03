package service

import (
	"context"
	"strconv"

	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/ininzzz/summer-backend/model"
	"github.com/sirupsen/logrus"
)

var BlogService blogService

type blogService struct {
	blogRepo infra.BlogRepo
	userRepo infra.UserRepo
}

func (u *blogService) List(ctx context.Context, blogListDTO *dto.BlogListRequestDTO) (*common.Response, error) {
	userID, err := strconv.Atoi(blogListDTO.UserID)
	if err != nil {
		logrus.Errorf("[blogService List] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	userID64 := int64(userID)
	blogs, err := u.blogRepo.Find(ctx, &infra.BlogQuery{
		UserID: &userID64,
	})
	if err != nil {
		logrus.Errorf("[blogService List] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := []*dto.BlogListResponseDTO{}
	for _, blog := range blogs {
		data = append(data, &dto.BlogListResponseDTO{
			ID:    blog.ID,
			Title: blog.Title,
			Like:  blog.Like,
		})
	}
	return common.NewResponseOfSuccess(data), nil
}

func (u *blogService) ListAll(ctx context.Context, blogListAllDTO *dto.BlogListAllRequestDTO) (*common.Response, error) {
	blogs, err := u.blogRepo.Find(ctx, &infra.BlogQuery{})
	if err != nil {
		logrus.Errorf("[blogService ListAll] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := []*dto.BlogListAllResponseDTO{}
	for _, blog := range blogs {
		data = append(data, &dto.BlogListAllResponseDTO{
			ID:     blog.ID,
			UserID: blog.UserID,
			Title:  blog.Title,
			Like:   blog.Like,
		})
	}
	return common.NewResponseOfSuccess(data), nil
}

func (u *blogService) Info(ctx context.Context, blogInfoDTO *dto.BlogInfoRequestDTO) (*common.Response, error) {
	blogID, err := strconv.Atoi(blogInfoDTO.BlogID)
	if err != nil {
		logrus.Errorf("[blogService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	blog, err := u.blogRepo.FindByID(ctx, int64(blogID))
	if err != nil {
		logrus.Errorf("[blogService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.BlogInfoResponseDTO{
		Title: blog.Title,
		Text:  blog.Text,
		Like:  blog.Like,
	}
	return common.NewResponseOfSuccess(data), nil
}

func (u *blogService) Post(ctx context.Context, blogPostDTO *dto.BlogPostRequestDTO) (*common.Response, error) {
	userID, err := strconv.Atoi(blogPostDTO.UserID)
	if err != nil {
		logrus.Errorf("[blogService Post] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	err = u.blogRepo.Save(ctx, &model.Blog{
		Title:  blogPostDTO.Title,
		UserID: int64(userID),
		Text:   blogPostDTO.Text,
	})
	if err != nil {
		logrus.Errorf("[blogService Post] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	return common.NewResponseOfSuccess(nil), nil
}

func (u *blogService) Like(ctx context.Context, blogLikeDTO *dto.BlogLikeRequestDTO) (*common.Response, error) {
	blog, err := u.blogRepo.FindByID(ctx, blogLikeDTO.BlogID)
	if err != nil {
		logrus.Errorf("[blogService Like] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	err = u.blogRepo.UpdateField(ctx, &model.Blog{
		ID:   blogLikeDTO.BlogID,
		Like: blog.Like + blogLikeDTO.Value,
	})
	if err != nil {
		logrus.Errorf("[blogService Like] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	return common.NewResponseOfSuccess(nil), nil
}

func (u *blogService) Comment(ctx context.Context, blogCommentDTO *dto.BlogCommentPostRequestDTO) (*common.Response, error) {
	userID, err := strconv.Atoi(blogCommentDTO.UserID)
	if err != nil {
		logrus.Errorf("[blogService Comment] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	user, err := u.userRepo.FindByID(ctx, int64(userID))
	if err != nil {
		logrus.Errorf("[blogService Comment] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	blog, err := u.blogRepo.FindByID(ctx, blogCommentDTO.BlogID)
	if err != nil {
		logrus.Errorf("[blogService Comment] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	blog.Comment = append(blog.Comment, dto.BlogCommentListResponseDTO{
		Username: user.Username,
		Comment:  blogCommentDTO.Comment,
	})
	err = u.blogRepo.Save(ctx, blog)
	if err != nil {
		logrus.Errorf("[blogService Comment] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	return common.NewResponseOfSuccess(nil), nil
}

func (u *blogService) CommentList(ctx context.Context, blogCommentListDTO *dto.BlogCommentListRequestDTO) (*common.Response, error) {
	blogID, err := strconv.Atoi(blogCommentListDTO.BlogID)
	if err != nil {
		logrus.Errorf("[blogService CommentList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	blog, err := u.blogRepo.FindByID(ctx, int64(blogID))
	if err != nil {
		logrus.Errorf("[blogService CommentList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := blog.Comment
	return common.NewResponseOfSuccess(data), nil
}
