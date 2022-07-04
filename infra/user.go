package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"gorm.io/gorm"
)

type User struct {
	ID         int64  `gorm:"primary_key"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	UserAvatar string `gorm:"column:user_avatar"`
	Gender     int    `gorm:"column:gender"`
	Email      string `gorm:"column:email"`
	Nickname   string `gorm:"column:nick_name"`
}

type UserQuery struct {
	ID       *int64
	Username *string
	Password *string
}

type UserRepo struct {
}

func (u *UserRepo) Save(ctx context.Context, user *model.User) error {
	userDO, err := u.toDO(user)
	if err != nil {
		return err
	}
	err = db.Where("id = ?", userDO.ID).Error
	if err == gorm.ErrRecordNotFound {
		err = db.Create(userDO).Error
	} else if err == nil {
		err = db.Save(userDO).Error
	}
	if err != nil {
		return err
	}
	return nil
}

//根据用户ID查找用户
func (u *UserRepo) FindByID(ctx context.Context, user *UserQuery) ([]*model.User, error) {
	userDOs := []*User{}
	if user.ID != nil {
		db = db.Where("username = ?", user.Username)
	}
	err := db.Find(&userDOs).Error
	if err != nil {
		return nil, err
	}
	ans := []*model.User{}
	for _, userDO := range userDOs {
		user, err := u.toModel(userDO)
		if err != nil {
			return nil, err
		}
		ans = append(ans, user)
	}
	return ans, nil
}

func (u *UserRepo) Find(ctx context.Context, user *UserQuery) ([]*model.User, error) {
	userDOs := []*User{}
	if user.Username != nil {
		db = db.Where("username = ?", user.Username)
	}
	if user.Password != nil {
		db = db.Where("password = ?", user.Password)
	}
	err := db.Find(&userDOs).Error
	if err != nil {
		return nil, err
	}
	ans := []*model.User{}
	for _, userDO := range userDOs {
		user, err := u.toModel(userDO)
		if err != nil {
			return nil, err
		}
		ans = append(ans, user)
	}
	return ans, nil
}

func (u *UserRepo) toDO(user *model.User) (*User, error) {
	return &User{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		UserAvatar: user.UserAvatar,
	}, nil
}

func (u *UserRepo) toModel(user *User) (*model.User, error) {
	return &model.User{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		UserAvatar: user.UserAvatar,
	}, nil
}
