package controllers

import (
	ArticleModel "blog/internal/modules/article/models"
	"blog/internal/modules/article/requests/articles"
	ArticleService "blog/internal/modules/article/services"
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/old"
	"blog/pkg/sessions"
	"fmt"
	"net/http"
	"strconv"

	ArticleResponse "blog/internal/modules/article/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	// Get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// Find the article from the database
	article, err := controller.articleService.Find(id)

	// If the article is not found, show error page
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// if article found, render article template
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})
}

// 获取创建文章页面
func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{"title": "Create article"})
}

func (controller *Controller) Store(c *gin.Context) {
	// validate the request
	var request articles.StoreRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	// Create the article
	article, err := controller.articleService.StoreAsUser(request, user)

	// Check if there is any error on the article creation
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}

// 编辑文章
func (controller *Controller) Edit(c *gin.Context) {
	//获取文章ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	//查找文章
	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	//获取当前用户
	currentUser := helpers.Auth(c)

	//检查当前用户是否是作者
	if currentUser.ID != article.UserID {
		html.Render(c, http.StatusForbidden, "templates/errors/html/403", gin.H{"title": "Forbidden", "message": "您没有权限修改此文章"})
		return
	}

	//渲染编辑页面
	html.Render(c, http.StatusOK, "modules/article/html/edit", gin.H{
		"article": article,
	})
}

// 更新文章
func (controller *Controller) Update(c *gin.Context) {
	//获取文章ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	//查找文章
	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	//获取当前用户
	currentUser := helpers.Auth(c)

	//检查当前用户是否是作者
	if currentUser.ID != article.UserID {
		html.Render(c, http.StatusForbidden, "templates/errors/html/403", gin.H{"title": "Forbidden", "message": "您没有权限更新此文章"})
		return
	}

	//验证请求
	var request articles.UpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d/edit", id))
		return
	}

	// 准备更新的文章数据
	articleModel := ArticleModel.Article{
		Model:   gorm.Model{ID: uint(id)},
		Title:   request.Title,
		Content: request.Content,
		UserID:  currentUser.ID,
	}

	// 更新文章
	updatedArticle, err := controller.articleService.Update(articleModel)
	if err != nil {
		errors.Init()
		errors.Add("error", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d/edit", id))
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", updatedArticle.ID))
}

// 删除文章
func (controller *Controller) Delete(c *gin.Context) {
	// 获取文章ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	// 获取当前用户
	currentUser := helpers.Auth(c)
	fmt.Printf("当前用户ID: %d\n", currentUser.ID)

	// 根据文章ID查找文章
	article, err := controller.articleService.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "该文章不存在"})
		return
	}
	fmt.Printf("查找文章成功: ID=%d, UserID=%d\n", article.ID, article.UserID)

	// 检查当前用户是否是文章的作者
	if currentUser.ID != article.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "抱歉，您没有权限删除此文章"})
		return
	}

	// 删除文章
	if err := controller.articleService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("删除失败：%v", err)})
		return
	}

}

// 搜索文章
func (controller *Controller) Search(c *gin.Context) {
	// 获取当前用户
	currentUser := helpers.Auth(c)

	// 获取搜索关键词
	keyword := c.Query("keyword")

	// 初始化空的文章列表
	articles := ArticleResponse.Articles{
		Data: []ArticleResponse.Article{},
	}

	// 执行搜索
	if keyword != "" {
		articles = controller.articleService.Search(keyword)
	}

	// 渲染搜索结果页面
	html.Render(c, http.StatusOK, "modules/article/html/search", gin.H{
		"keyword":     keyword,
		"articles":    articles,
		"currentUser": currentUser,
		"AUTH":        currentUser, // 当前用户信息
		"ERRORS":      map[string]string{},
		"old":         map[string]string{},
	})
}
