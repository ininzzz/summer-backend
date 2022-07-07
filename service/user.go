package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ininzzz/summer-backend/cache"
	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/ininzzz/summer-backend/model"
	"github.com/ininzzz/summer-backend/utils"
	"github.com/sirupsen/logrus"
)

var UserService userService

type userService struct {
	userRepo infra.UserRepo
}

//service-用户凭借用户名和密码登录
func (u *userService) Login(ctx context.Context, loginDTO *dto.LoginRequestDTO) (*common.Response, error) {
	users, err := u.userRepo.Find(ctx, &infra.UserQuery{
		Username: &loginDTO.Username,
		Password: &loginDTO.Password,
	})
	if err != nil {
		logrus.Errorf("[userService Login] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	} else if len(users) == 0 {
		err := fmt.Errorf("record not found")
		logrus.Errorf("[userService Login] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	token, err := common.GenerateToken(users[0].ID)
	if err != nil {
		logrus.Errorf("[userService Login] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.LoginResponseDTO{
		UserID: users[0].ID,
		Token:  token,
	}
	return common.NewResponseOfSuccess(data), nil
}

//service-注册
func (u *userService) Register(ctx context.Context, reqDTO *dto.User_Register_ReqDTO) (*common.Response, error) {
	//验证邮箱以及验证码是否在redis中有记录
	code, err := cache.RedisClient.Get(ctx, cache.EmailCodeKP+reqDTO.Email).Result()
	if err != nil { //redis取出邮箱验证码信息失败
		if err == cache.Redis_nil {
			data := &dto.User_Register_RespDTO{
				Ok:  false,
				Msg: "邮箱验证码信息不存在",
			}
			return common.NewResponseOfSuccess(data), nil
		}
		return common.NewResponseOfErr(err), err
	} else {
		if code != reqDTO.Verification { //验证未通过
			data := &dto.User_Register_RespDTO{
				Ok:  false,
				Msg: "邮箱验证码错误",
			}
			return common.NewResponseOfSuccess(data), nil
		} else { //验证通过
			//在mysql表中创建用户信息，返回UID
			user_info_resp, err := u.userRepo.CreateUser(ctx, &model.User{
				Username: reqDTO.Username,
				Password: reqDTO.Password,
				Email:    reqDTO.Email,
			})
			if user_info_resp != nil { //注册成功
				//用户信息
				user_info := dto.User_Register_UserInfo{
					UserID:   user_info_resp.ID,
					UserName: user_info_resp.Username,
					Email:    user_info_resp.Email,
					Avatar:   user_info_resp.UserAvatar,
				}
				//根据UID生成token
				token, _ := common.GenerateToken(user_info_resp.ID)
				data := &dto.User_Register_RespDTO{
					Ok:       true,
					Token:    token, //需要生成token
					UserInfo: user_info,
				}
				return common.NewResponseOfSuccess(data), nil
			} else { //注册失败
				if err == nil { //没有err，但注册失败（邮箱关联的账户已存在）
					data := &dto.User_Register_RespDTO{
						Ok:  false,
						Msg: "邮箱关联的账户已存在",
					}
					return common.NewResponseOfSuccess(data), err
				} //有err，注册失败
				return common.NewResponseOfErr(err), err
			}
		}
	}
}

//service-发送验证码到指定邮件地址
func (u *userService) EmailCode(ctx context.Context, reqDTO *dto.User_Email_Code_ReqDTO) (*common.Response, error) {
	email := reqDTO.Email
	//邮件地址不为空，邮箱地址有效性本应由前端用正则表达式校验来保证
	if email != "" {
		//生成随机数6位
		code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		//发送邮件，这一步有点耗时，可能需要优化(消息队列 或 先返回Ok=true再发邮件)
		err := utils.SendMail("无情的邮件发送人", []string{email}, "注册验证码", fmt.Sprintf("您正在尝试注册，这是您的验证码：%s，该验证码有效期为%s秒。", code, "300"))
		if err != nil {
			logrus.Errorf("[userService EmailCode] err: %v", err.Error())
			return common.NewResponseOfErr(err), err
		}
		//存储到redis中
		cache.RedisClient.Set(ctx, cache.EmailCodeKP+email, code, 300*time.Second)
		//返回是否成功的布尔值
		data := &dto.User_Email_Code_RespDTO{
			Ok: true,
		}
		return common.NewResponseOfSuccess(data), nil
	} else { //邮箱为空
		data := &dto.User_Email_Code_RespDTO{
			Ok: false,
		}
		return common.NewResponseOfSuccess(data), nil
	}
}

//service-根据用户ID获取用户信息
func (u *userService) FindInfoByID(ctx context.Context, infoDTO *dto.InfoRequestDTO) (*common.Response, error) {
	user, err := u.userRepo.FindByID(ctx, &infra.UserQuery{
		ID: &infoDTO.UserID,
	})
	if err != nil {
		logrus.Errorf("[userService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.InfoResponseDTO{
		Username: user.Username,
		Gender:   user.Gender,
		Email:    user.Email,
		Avatar:   user.UserAvatar,
	}
	return common.NewResponseOfSuccess(data), nil
}
