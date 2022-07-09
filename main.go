package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/cache"
	"github.com/ininzzz/summer-backend/tasks"
	"github.com/ininzzz/summer-backend/utils"
	"github.com/ininzzz/summer-backend/web"
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func InitConf() {
	// 从本地读取环境变量 .env文件
	godotenv.Load()
	// 连接Redis数据库
	cache.InitRedis()
	// 启动定时任务
	tasks.CronJob()
	//初始化邮件设置
	utils.Init_Email_Conf()
	//连接COS
	utils.CreateCOSClient()
}

//注册路由
func Register(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", web.UserWebHandler.Login)          //用户登录
		userGroup.PUT("/register", web.UserWebHandler.Register)     //用户注册
		userGroup.POST("/email/code", web.UserWebHandler.EmailCode) //请求发送邮箱验证码
		userGroup.Use(utils.JwtAuth())                              //需要登录的路由
		{
			userGroup.GET("/info", web.UserWebHandler.Info) //请求用户信息
		}
	}
	blogGroup := r.Group("/blog")
	{
		// blogGroup.POST("/post", web.BlogWebHandler.BlogPost)           //发布帖子
		blogGroup.GET("/home/list", web.BlogWebHandler.HomeList)       //查看首页帖子【基于滚动分页】
		blogGroup.GET("/space/list", web.BlogWebHandler.SpaceList)     //获取某个用户发布的所有帖子【不分页】
		blogGroup.GET("/info", web.BlogWebHandler.Info)                //获取某个帖子内容
		blogGroup.GET("/comment/list", web.BlogWebHandler.CommentList) //获取某个帖子的所有评论信息
		blogGroup.Use(utils.JwtAuth())                                 //需要登录的路由
		{
			blogGroup.PUT("/comment/post", web.BlogWebHandler.BlogCommentPost) //发布评论
			blogGroup.POST("/post", web.BlogWebHandler.BlogPost)               //发布帖子
		}
	}
}

//程序入口
func main() {
	InitConf()
	r := gin.Default()
	Register(r)
	r.Run("127.0.0.1:9090")
}
