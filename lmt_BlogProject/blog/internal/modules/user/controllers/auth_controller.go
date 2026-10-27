package controllers

import (
	"blog/internal/modules/user/helpers"
	"blog/internal/modules/user/information"
	"blog/internal/modules/user/requests/auth"
	UserService "blog/internal/modules/user/services"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/old"
	"blog/pkg/sessions"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate the request
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// after creating the user > redirect to home page
	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// validate the request
	var request auth.LoginRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// after creating the user > redirect to home page
	log.Printf("The user logged in successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")

	c.Redirect(http.StatusFound, "/")
}

// 查询个人信息
func (controller *Controller) Information(c *gin.Context) {
	// 获取当前登录用户
	currentUser := helpers.Auth(c)
	if currentUser.ID == 0 {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// 渲染个人信息页面
	html.Render(c, http.StatusOK, "modules/user/html/information", gin.H{
		"title": "个人信息",
		"user":  currentUser,
		"AUTH":  currentUser,
	})
}

// 编辑个人信息
func (controller *Controller) InformationEdit(c *gin.Context) {
	// 获取当前登录用户
	currentUser := helpers.Auth(c)
	if currentUser.ID == 0 {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// 渲染编辑页面
	html.Render(c, http.StatusOK, "modules/user/html/information_edit", gin.H{
		"user": currentUser,
		"AUTH": currentUser,
	})
}

// 更新个人信息
func (controller *Controller) InformationUpdate(c *gin.Context) {
	// 获取当前登录用户
	currentUser := helpers.Auth(c)
	if currentUser.ID == 0 {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// 验证更新请求
	var request information.UpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/information/edit")
		return
	}

	// 检查新邮箱是否已存在
	if request.Email != currentUser.Email && controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("email", "抱歉，该邮箱已被使用")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		c.Redirect(http.StatusFound, "/information/edit")
		return
	}

	// 更新用户信息
	_, err := controller.userService.UpdateInformation(currentUser.ID, request)
	if err != nil {
		errors.Init()
		errors.Add("error", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		c.Redirect(http.StatusFound, "/information/edit")
		return
	}

	// 返回信息页面，显示更新后的信息
	c.Redirect(http.StatusFound, "/information")
}
