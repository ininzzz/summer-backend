package main

import (
	"github.com/ininzzz/summer-backend/cache"
	"github.com/ininzzz/summer-backend/tasks"
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func InitConfigure() {
	// 从本地读取环境变量 .env文件
	godotenv.Load()

	// 连接Redis数据库
	cache.InitRedis()

	// 启动定时任务
	tasks.CronJob()
}
