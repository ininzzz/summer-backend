package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/service"
	"github.com/sirupsen/logrus"
)

var BlogWebHandler = &blogWebHandler{}

type blogWebHandler struct{}

func (u *blogWebHandler) List(c *gin.Context) {
	dto := dto.BlogListRequestDTO{
		UserID: c.Param("user_id"),
	}
	resp, err := service.BlogService.List(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler List] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (u *blogWebHandler) Info(c *gin.Context) {
	dto := dto.BlogInfoRequestDTO{
		BlogID: c.Param("blog_id"),
	}
	resp, err := service.BlogService.Info(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler Info] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
