package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/sirupsen/logrus"
)

func JwtAuth(ctx *gin.Context) {
	ans, err := common.ParseToken(strings.Split(ctx.GetHeader("Authorization"), " ")[1])
	if err != nil {
		logrus.Errorf("[register] err: %v", err.Error())
		return
	}
	ctx.Set("UserID", ans.UserID)
}
