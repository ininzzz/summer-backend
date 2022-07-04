package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/sirupsen/logrus"
)

var UserService userService

type userService struct {
	userRepo infra.UserRepo
}

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

func (u *userService) Info(ctx context.Context, infoDTO *dto.InfoRequestDTO) (*common.Response, error) {
	userID, err := strconv.Atoi(infoDTO.UserID)
	if err != nil {
		logrus.Errorf("[blogService BlogInfo] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	user, err := u.userRepo.FindByID(ctx, int64(userID))
	if err != nil {
		logrus.Errorf("[userService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.InfoResponseDTO{
		Username: user.Username,
		Gender:   user.Gender,
		Email:    user.Email,
		Icon:     string(user.Icon),
	}
	return common.NewResponseOfSuccess(data), nil
}

func (u *userService) FindInfoByID(ctx context.Context, infoDTO *dto.InfoRequestDTO) (*common.Response, error) {
	users, err := u.userRepo.FindByID(ctx, &infra.UserQuery{
		ID: &infoDTO.UserID,
	})
	if err != nil {
		logrus.Errorf("[userService Info] err: %v", err.Error())
		return common.NewResponseOfErr(err), err
	}
	data := &dto.InfoResponseDTO{
		Username: users[0].Username,
	}
	return common.NewResponseOfSuccess(data), nil
}
