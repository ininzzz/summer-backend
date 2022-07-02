package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type UserQuery struct {
	Username *string
	Password *string
}

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) Save(ctx context.Context, user *model.User) error {
	err := u.db.Where("id = ?", user.ID).Error
	if err == gorm.ErrRecordNotFound {
		err = u.db.Create(user).Error
	} else if err == nil {
		err = u.db.Save(user).Error
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) Find(ctx context.Context, user *UserQuery) ([]*model.User, error) {
	ans := []*model.User{}
	if user.Username != nil {
		u.db = u.db.Where("username = ?", user.Username)
	}
	if user.Password != nil {
		u.db = u.db.Where("password = ?", user.Password)
	}
	err := u.db.Find(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil
}

func (u *UserRepo) toDO(user *model.User) (*User, error) {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (u *UserRepo) toModel(user *User) (*model.User, error) {
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}
