package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Blog struct {
	ID     int64  `gorm:"primary_key"`
	UserID int64  `gorm:"column:user_id"`
	Title  string `gorm:"column:title"`
	Text   string `gorm:"column:text"`
}

type BlogQuery struct {
	ID     *int64
	UserID *int64
	Title  *string
}

type BlogRepo struct {
}

func (b *BlogRepo) Save(ctx context.Context, blog *model.Blog) error {
	BlogDO, err := b.toDO(blog)
	if err != nil {
		logrus.Errorf("[BlogRepo Save] err: %v", err.Error())
		return err
	}
	err = db.Where("id = ?", BlogDO.ID).Error
	if err == gorm.ErrRecordNotFound {
		err = db.Create(BlogDO).Error
	} else if err == nil {
		err = db.Save(BlogDO).Error
	}
	if err != nil {
		logrus.Errorf("[BlogRepo Save] err: %v", err.Error())
		return err
	}
	return nil
}

func (b *BlogRepo) Find(ctx context.Context, blog *BlogQuery) ([]*model.Blog, error) {
	blogDOs := []*Blog{}
	if blog.ID != nil {
		db = db.Where("id = ?", blog.ID)
	}
	if blog.UserID != nil {
		db = db.Where("user_id = ?", blog.UserID)
	}
	if blog.Title != nil {
		db = db.Where("title = ?", blog.Title)
	}
	err := db.Find(&blogDOs).Error
	if err != nil {
		logrus.Errorf("[BlogRepo Find] err: %v", err.Error())
		return nil, err
	}
	ans := []*model.Blog{}
	for _, blogDO := range blogDOs {
		blog, err := b.toModel(blogDO)
		if err != nil {
			logrus.Errorf("[BlogRepo Find] err: %v", err.Error())
			return nil, err
		}
		ans = append(ans, blog)
	}
	return ans, nil
}

func (b *BlogRepo) toDO(blog *model.Blog) (*Blog, error) {
	return &Blog{
		ID:   blog.ID,
		Text: blog.Text,
	}, nil
}

func (b *BlogRepo) toModel(blog *Blog) (*model.Blog, error) {
	return &model.Blog{
		ID:   blog.ID,
		Text: blog.Text,
	}, nil
}
