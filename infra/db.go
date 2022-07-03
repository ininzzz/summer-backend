package infra

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
