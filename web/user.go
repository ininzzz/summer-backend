package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/service"
	"github.com/sirupsen/logrus"
)

var UserWebHandler = &userWebHandler{}

type userWebHandler struct{}

func (u *userWebHandler) Login(c *gin.Context) {
	dto := dto.LoginRequestDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		logrus.Errorf("[userWebHandler Login] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
		return
	}
	resp, err := service.UserService.Login(c, &dto)
	if err != nil {
		logrus.Errorf("[userWebHandler Login] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (u *userWebHandler) Info(c *gin.Context) {
	dto := dto.InfoRequestDTO{
		UserID: c.Query("user_id"),
	}
	resp, err := service.UserService.Info(c, &dto)
	if err != nil {
		logrus.Errorf("[userWebHandler Info] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
