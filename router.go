package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/utils"
	"github.com/ininzzz/summer-backend/web"
)

func register(r *gin.Engine) {
	r.Use(utils.JwtAuth)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", web.UserWebHandler.Login)
		userGroup.GET("/info", web.UserWebHandler.Info)
	}
	blogGroup := r.Group("/blog")
	{

		blogGroup.GET("/home/list",web.BlogWebHandler.HomeList)  // 查看首页帖子【基于滚动分页】
		blogGroup.GET("/space/list",web.BlogWebHandler.SpaceList)  //获取某个用户发布的所有帖子【不分页】
		blogGroup.GET("/info", web.BlogWebHandler.Info)  //获取某个帖子内容
		blogGroup.GET("/comment/list", web.BlogWebHandler.CommentList)  //获取某个帖子的所有评论信息

	}
}
