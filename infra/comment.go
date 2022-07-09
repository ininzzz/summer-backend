package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"github.com/sirupsen/logrus"
)

//数据库中存储的blog表结构
type Comment struct {
	CommentID       int64  `gorm:"primary_key"`
	BlogID          int64  `gorm:"column:blog_id"`
	UserID          int64  `gorm:"column:user_id"`
	Text            string `gorm:"column:text"`
	CreateTimeStamp int64  `gorm:"column:create_time_stamp"`
}

//查询用
type CommentQuery struct {
	BlogID *int64
}

type CommentRepo struct {
}

//根据model中Comment的信息创建comment，失败返回nil，
func (repo *CommentRepo) CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	db := GetDB(ctx)
	//model转do
	CommentDO, err := repo.toDO(comment)
	if err != nil {
		return nil, err
	}
	err = db.Create(CommentDO).Error
	if err != nil {
		return nil, err
	}
	commentModel, _ := repo.toModel(CommentDO)
	return commentModel, nil
}

//根据BlogQuery中的参数查询,返回model类型的blog
func (repo *CommentRepo) Find(ctx context.Context, cmt *CommentQuery) ([]*model.Comment, error) {
	db := GetDB(ctx)
	commentDOs := []*Comment{}
	if cmt.BlogID != nil {
		db = db.Where("blog_id = ?", cmt.BlogID)
	}
	err := db.Find(&commentDOs).Error
	if err != nil {
		logrus.Errorf("[CommentRepo Find] err: %v", err.Error())
		return nil, err
	}
	ans := []*model.Comment{}
	for _, blogDO := range commentDOs {
		blog, err := repo.toModel(blogDO)
		if err != nil {
			logrus.Errorf("[CommentRepo Find] err: %v", err.Error())
			return nil, err
		}
		ans = append(ans, blog)
	}
	return ans, nil
}

func (repo *CommentRepo) toDO(cmt *model.Comment) (*Comment, error) {
	return &Comment{
		CommentID:       cmt.CommentID,
		BlogID:          cmt.BlogID,
		Text:            cmt.Text,
		UserID:          cmt.UserID,
		CreateTimeStamp: cmt.CreateTimeStamp,
	}, nil
}

func (repo *CommentRepo) toModel(cmt *Comment) (*model.Comment, error) {
	return &model.Comment{
		CommentID:       cmt.CommentID,
		BlogID:          cmt.BlogID,
		Text:            cmt.Text,
		UserID:          cmt.UserID,
		CreateTimeStamp: cmt.CreateTimeStamp,
	}, nil
}
