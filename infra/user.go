package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"gorm.io/gorm"
)

type User struct {
	ID         int64  `gorm:"column:user_id;primary_key"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	UserAvatar string `gorm:"column:user_avatar;default:default_avatar_url"`
	Gender     int    `gorm:"column:gender"`
	Email      string `gorm:"column:email"`
}

type UserQuery struct {
	ID       *int64
	Username *string
	Password *string
}

type UserRepo struct {
}

//根据model中User的信息（username+password+email）创建用户，失败返回nil，
func (u *UserRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	db := GetDB(ctx)
	//model转do
	userDO, err := u.toDO(user)
	if err != nil {
		return nil, err
	}
	//查看邮件是否已存在
	var resDO *User
	err = db.Where("email = ?", userDO.Email).First(&resDO).Error
	if err == gorm.ErrRecordNotFound { //邮件不存在，正常注册
		err = db.Create(userDO).Error
		if err != nil {
			return nil, err
		}
		userModel, _ := u.toModel(userDO)
		return userModel, nil
	} else if err == nil { //邮件已存在,仍然返回错误信息
		return nil, err
	}
	return nil, err //其他错误
}

//
func (u *UserRepo) Save(ctx context.Context, user *model.User) error {
	db := GetDB(ctx)
	userDO, err := u.toDO(user)
	if err != nil {
		return err
	}
	err = db.Where("user_id = ?", userDO.ID).Error
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

func (u *UserRepo) FindByID(ctx context.Context, user *UserQuery) (*model.User, error) {
	db := GetDB(ctx)
	userDO := User{}
	err := db.Where("user_id = ?", user.ID).Find(&userDO).Error
	if err != nil {
		return nil, err
	}
	res, err := u.toModel(&userDO)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserRepo) Find(ctx context.Context, user *UserQuery) ([]*model.User, error) {
	db := GetDB(ctx)
	userDOs := []*User{}
	if user.ID != nil {
		db = db.Where("user_id = ?", user.ID)
	}
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
	gender := 0
	if user.Gender == "女" {
		gender = 1
	}
	return &User{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		UserAvatar: user.UserAvatar,
		Gender:     gender,
		Email:      user.Email,
	}, nil
}

func (u *UserRepo) toModel(user *User) (*model.User, error) {
	gender := ""
	if user.Gender == 0 {
		gender = "男"
	} else {
		gender = "女"
	}
	return &model.User{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		UserAvatar: user.UserAvatar,
		Gender:     gender,
		Email:      user.Email,
	}, nil
}
