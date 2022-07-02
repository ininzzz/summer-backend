package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/web"
)

func register(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", web.UserWebHandler.Login)
	}
}
