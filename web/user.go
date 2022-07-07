package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/service"
	"github.com/sirupsen/logrus"
)

var UserWebHandler = &userWebHandler{}

type userWebHandler struct{}

// 登录
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

// 注册
func (u *userWebHandler) Register(c *gin.Context) {
	dto := dto.User_Register_ReqDTO{}
	//绑定
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		logrus.Errorf("[userWebHandler Register bind] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
		return
	}
	resp, err := service.UserService.Register(c, &dto)
	if err != nil {
		logrus.Errorf("[userWebHandler Register service] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// 请求邮箱验证码
func (u *userWebHandler) EmailCode(c *gin.Context) {
	dto := dto.User_Email_Code_ReqDTO{}
	//绑定
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		logrus.Errorf("[userWebHandler EmailCode] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
		return
	}
	//调用service
	resp, err := service.UserService.EmailCode(c, &dto)
	if err != nil {
		logrus.Errorf("[userWebHandler EmailCode] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//请求用户信息
func (u *userWebHandler) Info(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	dto := dto.InfoRequestDTO{
		UserID: user_id,
	}
	resp, err := service.UserService.FindInfoByID(c, &dto)
	if err != nil {
		logrus.Errorf("[userWebHandler Info] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
