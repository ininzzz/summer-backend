package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//数据库中存储的blog表结构
type Comment struct {
	BlogID          int64  `gorm:"primary_key"`
	UserID          int64  `gorm:"column:user_id"`
	Text            string `gorm:"column:text"`
	CreateTimestamp int64  `gorm:"column:create_time_stamp"`
	ModifyTimestamp int64  `gorm:"column:modify_time_stamp"`
}

//查询用
type CommentQuery struct {
	BlogID *int64
}

type CommentRepo struct {
}

//转换model中的blog类型为数据库中的Blog，存储之
func (repo *CommentRepo) Save(ctx context.Context, cmt *model.Comment) error {
	CmtDO, err := repo.toDO(cmt)
	if err != nil {
		logrus.Errorf("[CommentRepo Save] err: %v", err.Error())
		return err
	}
	err = db.Where("id = ?", CmtDO.BlogID).Error
	if err == gorm.ErrRecordNotFound {
		err = db.Create(CmtDO).Error
	} else if err == nil {
		err = db.Save(CmtDO).Error
	}
	if err != nil {
		logrus.Errorf("[CommentRepo Save] err: %v", err.Error())
		return err
	}
	return nil
}

//根据BlogQuery中的参数查询,返回model类型的blog
func (repo *CommentRepo) Find(ctx context.Context, cmt *CommentQuery) ([]*model.Comment, error) {
	commentDOs := []*Comment{}
	if cmt.BlogID != nil {
		db = db.Where("id = ?", cmt.BlogID)
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
		BlogID: cmt.BlogID,
		Text:   cmt.Text,
		UserID: cmt.UserID,
	}, nil
}

func (repo *CommentRepo) toModel(cmt *Comment) (*model.Comment, error) {
	return &model.Comment{
		BlogID: cmt.BlogID,
		Text:   cmt.Text,
		UserID: cmt.UserID,
	}, nil
}
