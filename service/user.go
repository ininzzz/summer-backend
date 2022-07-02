package service

import (
	"context"

	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/infra"
	"github.com/ininzzz/summer-backend/model"
)

var UserService userService

type userService struct {
	repo infra.UserRepo
}

func (u *userService) Login(ctx context.Context, loginDTO dto.LoginRequestDTO) (*common.Response, error) {
	user := model.User{
		Username: loginDTO.Username,
		Password: loginDTO.Password,
	}
	_, err := u.repo.Find(ctx, &infra.UserQuery{
		Username: &user.Username,
	})
	if err != nil {
		return common.NewResponseOfErr(err), err
	}
	data := dto.LoginResponseDTO{
		Token: "Bearer xxx",
	}
	return common.NewResponseOfSuccess(data), nil
}
