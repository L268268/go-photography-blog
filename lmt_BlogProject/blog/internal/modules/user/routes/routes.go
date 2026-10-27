package routes

import (
	"blog/internal/middlewares"
	userCtrl "blog/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", userController.HandleLogout)
		// 定义一个 GET 请求的路由 "/information"，客户端发起 GET 请求访问此路径
		// 调用 userController.Information 函数来处理，该函数负责获取并返回用户的个人信息页面给客户端，比如渲染包含用户信息的 HTML 页面等操作。
		authGroup.GET("/information", userController.Information)

		// 定义一个 GET 请求的路由 "/information/edit"，客户端通过 GET 请求访问该路径，
		// 调用 userController.InformationEdit 函数进行处理，此函数用于渲染用户编辑个人信息的页面，方便用户对自己的信息进行修改编辑操作。
		authGroup.GET("/information/edit", userController.InformationEdit)

		// 定义一个 POST 请求的路由 "/information/update"，客户端发起 POST 请求到该路径，
		// 调用 userController.InformationUpdate 函数来处理这个请求，这个函数实现了更新用户个人信息的具体逻辑，比如将用户提交修改后的信息保存到数据库等操作。
		authGroup.POST("/information/update", userController.InformationUpdate)
	}
}
