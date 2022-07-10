package service

import (
	"context"
	"time"

	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/ininzzz/summer-backend/model"
	"github.com/sirupsen/logrus"
)

var BlogService blogService

type blogService struct {
	blogRepo    infra.BlogRepo
	commentRepo infra.CommentRepo
	likeRepo    infra.LikeRepo
}

//string数组连成一个string，用分号分割
func arrayToString(arr []string) string {
	var result string
	for _, i := range arr { //遍历数组中所有元素追加成string
		result += i
		result += ";"
	}
	return result
}

//service-发布blog
func (s *blogService) BlogPost(ctx context.Context, reqDTO *dto.Blog_Post_ReqDTO) (*common.Response, error) {
	//在mysql表中创建新的blog，返回blog_id
	blog_create_resp, err := s.blogRepo.CreateBlog(ctx, &model.Blog{
		UserID:          reqDTO.UserID,
		Text:            reqDTO.Text,
		Imgs:            arrayToString(reqDTO.Files),
		CreateTimestamp: time.Now().Unix(),
	})
	if err != nil {
		return common.NewResponseOfErr(err), err
	}
	//创建blog成功
	if blog_create_resp != nil {
		data := &dto.Blog_Post_RespDTO{
			Ok:      true,
			Message: "",
			BlogID:  blog_create_resp.BlogID,
		}
		return common.NewResponseOfSuccess(data), nil
	}
	return common.NewResponseOfSuccess(nil), nil
}

// blog/home/list
func (s *blogService) HomeList(ctx context.Context, reqDTO *dto.BlogHomeListRequestDTO) (*common.Response, error) {
	//根据发送的时间戳查询固定条数(比如10条)的blog条
	blogs, err := s.blogRepo.FindByTimeStamp(ctx, &infra.BlogQuery{ //此处查询本应返回最小时间戳填入下方LastTimeStamp
		CreateTimeStamp: &reqDTO.LastTimeStamp,
	})
	if err != nil {
		logrus.Errorf("[blogService BlogHomeList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	if len(blogs) != 0 {
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
				UserAvatar: user.UserAvatar,
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
	return common.NewResponseOfSuccess(nil), nil
}

// /blog/space 不分页
func (s *blogService) SpaceList(ctx context.Context, reqDTO *dto.BlogSpaceListRequestDTO) (*common.Response, error) {
	//根据user_id查找blog
	blogs, err := s.blogRepo.FindByUserID(ctx, &infra.BlogQuery{
		UserID: &reqDTO.UserID,
	})
	//出错返回
	if err != nil {
		logrus.Errorf("[blogService SpaceList] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	//查找结果不为空继续
	if len(blogs) != 0 {
		data := []*dto.BlogSpaceListResponseDTO{}
		for _, blog := range blogs {
			//对每个blog获取user信息
			user, _ := UserService.userRepo.FindByID(ctx, &infra.UserQuery{
				ID: &blog.UserID,
			})
			data = append(data, &dto.BlogSpaceListResponseDTO{
				BlogID:     blog.BlogID,
				UserID:     blog.UserID,
				UserName:   user.Username,
				UserAvatar: user.UserAvatar,
				Text:       blog.Text,
				Imgs:       blog.Imgs,
				Like:       blog.Like,
			})
		}
		return common.NewResponseOfSuccess(data), nil
	}
	//查找结果为空
	return common.NewResponseOfSuccess(nil), nil
}

// /blog/info
func (s *blogService) Info(ctx context.Context, reqDTO *dto.BlogInfoRequestDTO) (*common.Response, error) {
	//根据blog_id查找blog
	blogs, err := s.blogRepo.FindByBlogID(ctx, &infra.BlogQuery{
		BlogID: &reqDTO.BlogID,
	})
	//出错返回
	if err != nil {
		logrus.Errorf("[blogService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	//查找结果不为空继续
	if len(blogs) != 0 {
		//查找用户信息
		user, _ := UserService.userRepo.FindByID(ctx, &infra.UserQuery{
			ID: &blogs[0].UserID,
		})
		data := &dto.BlogInfoResponseDTO{
			UserID:     blogs[0].UserID,
			UserName:   user.Username,
			UserAvatar: user.UserAvatar,
			Imgs:       blogs[0].Imgs,
			Text:       blogs[0].Text,
			Like:       blogs[0].Like,
		}
		return common.NewResponseOfSuccess(data), nil
	}
	//查找结果为空
	return common.NewResponseOfSuccess(nil), nil
}

//service-发布comment
func (s *blogService) BlogCommentPost(ctx context.Context, reqDTO *dto.Blog_Comment_Post_ReqDTO) (*common.Response, error) {
	//在mysql表中创建新的comment
	comment_create_resp, err := s.commentRepo.CreateComment(ctx, &model.Comment{
		UserID:          reqDTO.UserID,
		Text:            reqDTO.Text,
		BlogID:          reqDTO.BlogID,
		CreateTimeStamp: time.Now().Unix(),
	})
	if err != nil {
		return common.NewResponseOfErr(err), err
	}
	//创建comment成功
	if comment_create_resp != nil {
		data := &dto.Blog_Comment_Post_RespDTO{
			Ok:      true,
			Message: "",
		}
		return common.NewResponseOfSuccess(data), nil
	}
	return common.NewResponseOfSuccess(nil), nil
}

//service-查询是否点赞了某条blog
func (s *blogService) BlogIfLiked(ctx context.Context, reqDTO *dto.Blog_If_Liked_ReqDTO) (*common.Response, error) {
	if reqDTO.UserID == 0 { //未登录
		data := &dto.Blog_If_Liked_RespDTO{
			Liked: false,
		}
		return common.NewResponseOfSuccess(data), nil
	}
	//查询是否存在
	resp, err := s.likeRepo.FindIfExist(ctx, &infra.LikeQuery{
		UserID: &reqDTO.UserID,
		BlogID: &reqDTO.BlogID,
	})
	if err != nil {
		return common.NewResponseOfErr(err), err
	}
	//查询成功
	data := &dto.Blog_If_Liked_RespDTO{
		Liked: resp,
	}
	return common.NewResponseOfSuccess(data), nil
}

//service-点赞/取消点赞某条blog
func (s *blogService) BlogLike(ctx context.Context, reqDTO *dto.Blog_Like_ReqDTO) (*common.Response, error) {
	//查询是否存在
	resp, err := s.likeRepo.AddOrRemove(ctx, &infra.LikeQuery{
		UserID: &reqDTO.UserID,
		BlogID: &reqDTO.BlogID,
	})
	if err != nil {
		return common.NewResponseOfErr(err), err
	}
	//查询成功
	data := &dto.Blog_Like_RespDTO{
		Ok: resp,
	}
	return common.NewResponseOfSuccess(data), nil
}
