package controllers

import (
	"blog/internal/modules/comment/requests/comments"
	CommentService "blog/internal/modules/comment/services"
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/sessions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller 评论控制器
type Controller struct {
	commentService CommentService.CommentServiceInterface
}

// New 创建评论控制器
func New() *Controller {
	return &Controller{
		commentService: CommentService.New(),
	}
}

// CreateComment 创建评论
func (controller *Controller) Comment(c *gin.Context) {
	// 获取当前登录用户
	currentUser := helpers.Auth(c)

	// 验证请求数据
	var request comments.StoreRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(int(request.ArticleID)))
		return
	}

	// 创建评论
	err := controller.commentService.Create(request, currentUser)
	if err != nil {
		errors.Init()
		errors.Add("error", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(int(request.ArticleID)))
		return
	}

	// 创建成功，返回文章页面
	c.Redirect(http.StatusFound, "/articles/"+strconv.Itoa(int(request.ArticleID)))
}
