package service

import (
	"context"

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


// blog/home/list
func (u *blogService) HomeList(ctx context.Context, reqDTO *dto.BlogHomeListRequestDTO) (*common.Response, error) {
	//根据发送的时间戳查询固定条数(比如10条)的blog条
	blogs, err := u.blogRepo.FindByTimeStamp(ctx, &infra.BlogQuery{ //此处查询本应返回最小时间戳填入下方LastTimeStamp
		CreateTimeStamp: &reqDTO.LastTimeStamp,
	})
	if err != nil {
		logrus.Errorf("[blogService BlogHomeList] err: %v", err.Error())

		return common.NewResponseOfErr(err), err
	}
	//构造DTO中的BlogList
	blog_data := []dto.HomeListBlog{}
	for _, blog := range blogs {

		//对每条blog获取其他相关信息，比如UserAvatar
		user, _ := UserService.userRepo.FindByID(ctx, &infra.UserQuery{
			ID: &blog.UserID,
		})
		blog_data = append(blog_data, dto.HomeListBlog{
			BlogID:     blog.BlogID,
			UserID:     blog.UserID,
			UserAvatar: user[0].UserAvatar,
			Text:       blog.Text,
			//还没填完

		})
	}
	//构造DTO
	data := &dto.BlogHomeListResponseDTO{
		//LastTimeStamp: xxx, //填入本次查询到的最小的时间戳
		BlogList: blog_data,
	}
	return common.NewResponseOfSuccess(data), nil
}


// /blog/space 不分页
func (u *blogService) SpaceList(ctx context.Context, reqDTO *dto.BlogSpaceListRequestDTO) (*common.Response, error) {
	blogs, err := u.blogRepo.Find(ctx, &infra.BlogQuery{
		UserID: &reqDTO.UserID,
	})
	if err != nil {
		logrus.Errorf("[blogService SpaceList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := []*dto.BlogSpaceListResponseDTO{}
	for _, blog := range blogs {
		user, _ := UserService.userRepo.FindByID(ctx, &infra.UserQuery{
			ID: &blog.UserID,
		})
		data = append(data, &dto.BlogSpaceListResponseDTO{
			BlogID:     blog.BlogID,
			UserID:     blog.UserID,
			UserAvatar: user[0].UserAvatar,
			Text:       blog.Text,
			//还没填完
		})
	}
	return common.NewResponseOfSuccess(data), nil
}

// /blog/info
func (u *blogService) Info(ctx context.Context, reqDTO *dto.BlogInfoRequestDTO) (*common.Response, error) {
	blogs, err := u.blogRepo.Find(ctx, &infra.BlogQuery{
		BlogID: &reqDTO.BlogID,
	})

	if err != nil {
		logrus.Errorf("[blogService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.BlogInfoResponseDTO{

		Text: blogs[0].Text,
		//还没填完
	}
	return common.NewResponseOfSuccess(data), nil
}
