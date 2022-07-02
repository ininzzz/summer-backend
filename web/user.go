package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/service"
)

var UserWebHandler = &userWebHandler{}

type userWebHandler struct{}

func (u *userWebHandler) Login(c *gin.Context) {
	dto := dto.LoginRequestDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
	}
	resp, err := service.UserService.Login(c, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
	}
	c.JSON(http.StatusOK, resp)
}
