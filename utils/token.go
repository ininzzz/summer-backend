package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/sirupsen/logrus"
)

// func JwtAuth(ctx *gin.Context) {
// 	ans, err := common.ParseToken(strings.Split(ctx.GetHeader("Authorization"), " ")[1])
// 	if err != nil {
// 		logrus.Errorf("[register] err: %v", err.Error())
// 		return
// 	}
// 	fmt.Printf("ans: %v\n", ans)
// 	ctx.Set("UserID", ans.UserID)
// }

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ans, err := common.ParseToken(strings.Split(ctx.GetHeader("Authorization"), " ")[1])
		if err != nil {
			logrus.Errorf("[register] err: %v", err.Error())
			return
		}
		fmt.Printf("ans: %v\n", ans)
		ctx.Set("UserID", ans.UserID)
		ctx.Next()
	}
}
