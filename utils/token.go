package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/sirupsen/logrus"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Token := strings.Split(ctx.GetHeader("Authorization"), " ")
		if len(Token) < 2 {
			ctx.JSON(http.StatusBadRequest, "Authorization is invalid")
			ctx.Abort()
			return
		}
		ans, err := common.ParseToken(Token[1])
		if err != nil {
			logrus.Errorf("[jwt] err: %v", err.Error())
			ctx.JSON(http.StatusBadRequest, "Authorization is invalid")
			ctx.Abort()
			return
		}
		ctx.Set("UserID", ans.UserID)
		ctx.Next()
	}
}
