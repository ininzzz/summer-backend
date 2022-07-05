package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/service"
	"github.com/sirupsen/logrus"
)

var BlogWebHandler = &blogWebHandler{}

type blogWebHandler struct{}

// 查看首页帖子【基于滚动分页】  /home/list
func (u *blogWebHandler) HomeList(c *gin.Context) {
	lastTimeStamp, _ := strconv.ParseInt(c.Query("lastTimeStamp"), 10, 64)
	dto := dto.BlogHomeListRequestDTO{
		LastTimeStamp: lastTimeStamp,
	}
	resp, err := service.BlogService.HomeList(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler HomeList] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//获取某个用户发布的所有帖子【不分页】  /space/list
func (u *blogWebHandler) SpaceList(c *gin.Context) {
	UID, _ := strconv.ParseInt(c.Query("UserID"), 10, 64)
	dto := dto.BlogSpaceListRequestDTO{
		UserID: UID,
	}
	resp, err := service.BlogService.SpaceList(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler List] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//获取某个帖子下的所有评论信息【不分页】  /comment/list
func (u *blogWebHandler) CommentList(c *gin.Context) {
	BID, _ := strconv.ParseInt(c.Query("blog_id"), 10, 64)
	dto := dto.BlogCommentListRequestDTO{
		BlogID: BID,
	}
	resp, err := service.CommentService.CommentList(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler CommentList] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// blog/info
func (u *blogWebHandler) Info(c *gin.Context) {
	BID, _ := strconv.ParseInt(c.Query("blog_id"), 10, 64)
	dto := dto.BlogInfoRequestDTO{
		BlogID: BID,
	}
	resp, err := service.BlogService.Info(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler Info] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
