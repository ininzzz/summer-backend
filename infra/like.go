package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"gorm.io/gorm"
)

//数据库中存储的like表结构
type Like struct {
	BlogID int64 `gorm:"column:blog_id"`
	UserID int64 `gorm:"column:user_id"`
}

//查询用
type LikeQuery struct {
	BlogID *int64
	UserID *int64
}

type LikeRepo struct {
}

func (repo *LikeRepo) CreateLike(ctx context.Context, like *model.Like) (*model.Like, error) {
	db := GetDB(ctx)
	//model转do
	LikeDO, err := repo.toDO(like)
	if err != nil {
		return nil, err
	}
	err = db.Create(LikeDO).Error
	if err != nil {
		return nil, err
	}
	LikeModel, _ := repo.toModel(LikeDO)
	return LikeModel, nil
}

//查看是否存在某条记录
func (repo *LikeRepo) FindIfExist(ctx context.Context, like_query *LikeQuery) (bool, error) {
	db := GetDB(ctx)
	var like Like
	results := db.Table("like").Where("user_id = ? and blog_id = ?", like_query.UserID, like_query.BlogID).First(&like)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound { //没找到
			return false, nil
		}
		return false, results.Error //其他错误
	} else { //找到了
		return true, nil
	}
}

//查看是否存在某条记录，若存在则删除，若不存在则添加,返回是否执行成功
func (repo *LikeRepo) AddOrRemove(ctx context.Context, like_query *LikeQuery) (bool, error) {
	db := GetDB(ctx)
	var like Like
	results := db.Table("like").Where("user_id = ? and blog_id = ?", like_query.UserID, like_query.BlogID).First(&like)
	if results.Error != nil {
		if results.Error == gorm.ErrRecordNotFound { //没找到
			//add
			like := Like{UserID: *like_query.UserID, BlogID: *like_query.BlogID}
			results := db.Table("like").Create(&like)
			if results.Error != nil {
				return false, results.Error
			}
			return true, nil
		}
		return false, results.Error //其他错误
	} else { //找到了
		//remove
		res := results.Table("like").Delete(&like)
		if res.Error != nil {
			return false, res.Error
		}
		return true, nil
	}
}

func (repo *LikeRepo) toDO(like *model.Like) (*Like, error) {
	return &Like{
		BlogID: like.BlogID,
		UserID: like.UserID,
	}, nil
}

func (repo *LikeRepo) toModel(like *Like) (*model.Like, error) {
	return &model.Like{
		BlogID: like.BlogID,
		UserID: like.UserID,
	}, nil
}
