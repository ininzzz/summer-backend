package infra

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var globalDB *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	globalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logrus.Errorf("[db init] err: %v", err.Error())
		panic(err)
	}
	logrus.Infof("db init successfully")
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}
