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
		blogGroup.GET("/list/:user_id", web.BlogWebHandler.List)
		blogGroup.GET("/:blog_id", web.BlogWebHandler.Info)
	}
}
