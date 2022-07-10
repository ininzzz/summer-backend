package web

import (
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ininzzz/summer-backend/common"
	"github.com/ininzzz/summer-backend/dto"
	"github.com/ininzzz/summer-backend/service"
	"github.com/ininzzz/summer-backend/utils"
	"github.com/sirupsen/logrus"
)

var BlogWebHandler = &blogWebHandler{}

type blogWebHandler struct{}

// 查看首页帖子【基于滚动分页】  /home/list
func (w *blogWebHandler) HomeList(c *gin.Context) {
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
func (w *blogWebHandler) SpaceList(c *gin.Context) {
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
func (w *blogWebHandler) CommentList(c *gin.Context) {
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
func (w *blogWebHandler) Info(c *gin.Context) {
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

//发布帖子  /blog/post
func (w *blogWebHandler) BlogPost(c *gin.Context) {
	//绑定数据
	form, _ := c.MultipartForm()
	files := form.File["files"]
	text := form.Value["text"]
	file_url_list := []string{}
	// 遍历所有图片
	var url string
	for _, file := range files {
		//获取图像后缀
		suffix := path.Ext(file.Filename)
		//获取时间戳防止重复 !需要精准到纳秒，防止传输过快产生同名，然后出错
		time_stamp := time.Now().UnixNano()
		//获取一个1w以内的随机数
		rand_num := rand.Intn(10000)
		//将时间辍的类型转换
		format_time_stamp := strconv.FormatInt(time_stamp, 10)
		//将随机数转换类型
		format_rand_num := strconv.FormatInt(int64(rand_num), 10)
		name := format_time_stamp + format_rand_num + suffix
		// //写入保存位置与自定义名称，并且带上文件自带后缀名
		// dst := path.Join("./imgs", name)
		// // 存储文件
		// _ = c.SaveUploadedFile(file, dst)

		//上传cos
		f, _ := file.Open()
		err := utils.UploadImg_Reader(name, f) //通过io.Reader上传
		//err := utils.UploadImg_Local(name, dst) //先存储到本地再上传
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		//url添加至file列表
		url = utils.Img_URL + name
		file_url_list = append(file_url_list, url)
	}
	//绑定dto
	dto := dto.Blog_Post_ReqDTO{
		Files:  file_url_list,
		Text:   text[0],
		UserID: c.GetInt64("UserID"), //需要登录验证，不然会取到用户ID为0
	}
	//请求服务
	resp, err := service.BlogService.BlogPost(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogPost] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//发表评论 blog/comment/post
func (w *blogWebHandler) BlogCommentPost(c *gin.Context) {
	//绑定数据
	dto := dto.Blog_Comment_Post_ReqDTO{
		UserID: c.GetInt64("UserID"), //需要登录验证，不然会取到用户ID为0
	}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogCommentPost bind] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
		return
	}
	//请求服务
	resp, err := service.BlogService.BlogCommentPost(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogCommentPost] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//是否点赞过某条blog  blog/if/liked
func (w *blogWebHandler) BlogIfLiked(c *gin.Context) {
	//绑定数据
	dto := dto.Blog_If_Liked_ReqDTO{
		UserID: c.GetInt64("UserID"), //不需要登录验证，若没有登录会取到用户ID为0
	}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogIfLiked bind] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
		return
	}
	//请求服务
	resp, err := service.BlogService.BlogIfLiked(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogIfLiked] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//点赞/取消点赞某条blog  blog/like
func (w *blogWebHandler) BlogLike(c *gin.Context) {
	//绑定数据
	dto := dto.Blog_Like_ReqDTO{
		UserID: c.GetInt64("UserID"), //需要登录验证
	}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogLike bind] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, common.NewResponseOfErr(err))
		return
	}
	//请求服务
	resp, err := service.BlogService.BlogLike(c, &dto)
	if err != nil {
		logrus.Errorf("[blogWebHandler BlogLike service] err: %v", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
