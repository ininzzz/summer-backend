package infra

import (
	"context"

	"github.com/ininzzz/summer-backend/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//数据库中存储的blog表结构
type Blog struct {
	BlogID          int64  `gorm:"column:blog_id;primary_key"`
	UserID          int64  `gorm:"column:user_id"`
	Text            string `gorm:"column:text"`
	Imgs            string `gorm:"column:imgs"`
	CreateTimestamp int64  `gorm:"column:create_time_stamp"`
	ModifyTimestamp int64  `gorm:"column:modify_time_stamp"`
	Like            int    `gorm:"column:like"`
}

//查询用
type BlogQuery struct {
	BlogID          *int64
	UserID          *int64
	CreateTimeStamp *int64
}

type BlogRepo struct {
}

//转换model中的blog类型为数据库中的Blog，存储之
func (b *BlogRepo) Save(ctx context.Context, blog *model.Blog) error {
	db := GetDB(ctx)
	BlogDO, err := b.toDO(blog)
	if err != nil {
		logrus.Errorf("[BlogRepo Save] err: %v", err.Error())
		return err
	}
	err = db.Where("blog_id = ?", BlogDO.BlogID).Error
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

//根据生成时间戳查询
func (b *BlogRepo) FindByTimeStamp(ctx context.Context, blog *BlogQuery) ([]*model.Blog, error) {
	db := GetDB(ctx)
	blogDOs := []*Blog{}
	if blog.CreateTimeStamp != nil {
		//需要查询固定条数的比blog.CreateTimeStamp小的blog（且在满足条件的blog中有最大的时间戳），并跳过上次查询与最小时间戳相同的offset条blog，使用小于等于号检索（create_time_stamp <= ?）
		//逻辑：上次看到的最后一条blog的时间戳是blog.CreateTimeStamp,且与这条时间戳相同的有offset条，需要跳过这几条
		//简化后的逻辑：不考虑有两条blog拥有同样的时间戳，不需要考虑offset参数，使用小于号检索
		err := db.Where("create_time_stamp < ?", blog.CreateTimeStamp).Find(&blogDOs).Error //此处仅举个例子，具体用法我不太清楚
		if err != nil {
			logrus.Errorf("[BlogRepo Find] err: %v", err.Error())
			return nil, err
		}
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

func (b *BlogRepo) FindByUserID(ctx context.Context, blog *BlogQuery) ([]*model.Blog, error) {
	db := GetDB(ctx)
	blogDOs := []*Blog{}
	if blog.UserID != nil {
		err := db.Where("user_id = ?", blog.UserID).Find(&blogDOs).Error
		if err != nil {
			logrus.Errorf("[BlogRepo Find] err: %v", err.Error())
			return nil, err
		}
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

func (b *BlogRepo) FindByBlogID(ctx context.Context, blog *BlogQuery) ([]*model.Blog, error) {
	db := GetDB(ctx)
	blogDOs := []*Blog{}
	if blog.BlogID != nil {
		err := db.Where("blog_id = ?", blog.BlogID).Find(&blogDOs).Error
		if err != nil {
			logrus.Errorf("[BlogRepo Find] err: %v", err.Error())
			return nil, err
		}
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

func (b *BlogRepo) UpdateField(ctx context.Context, blog *model.Blog) error {
	db := GetDB(ctx)
	err := db.Model(&blog).Updates(blog).Error
	if err != nil {
		logrus.Errorf("[BlogRepo UpdateField] err: %v", err.Error())
		return err
	}
	return nil
}

func (b *BlogRepo) toDO(blog *model.Blog) (*Blog, error) {
	return &Blog{
		BlogID:          blog.BlogID,
		UserID:          blog.UserID,
		Text:            blog.Text,
		Imgs:            blog.Imgs,
		CreateTimestamp: blog.CreateTimestamp,
		ModifyTimestamp: blog.ModifyTimestamp,
	}, nil
}

func (b *BlogRepo) toModel(blog *Blog) (*model.Blog, error) {
	return &model.Blog{
		BlogID:          blog.BlogID,
		Text:            blog.Text,
		UserID:          blog.UserID,
		Imgs:            blog.Imgs,
		CreateTimestamp: blog.CreateTimestamp,
		ModifyTimestamp: blog.ModifyTimestamp,
	}, nil
}
